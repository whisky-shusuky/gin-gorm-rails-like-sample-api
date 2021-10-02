package test

import (
	"fmt"
	"gin-gorm-rails-like-sample-api/db"
	"gin-gorm-rails-like-sample-api/server"
	"gin-gorm-rails-like-sample-api/test/helper"
	test_db "gin-gorm-rails-like-sample-api/test/testdata/db"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	fixtures *testfixtures.Loader
)

func ExecTest(m *testing.M, testdb *gorm.DB) int {
	db.Db = testdb
	code := m.Run()
	return code
}

func TestMain(m *testing.M) {
	// Making test db
	testdb, pool := test_db.Init()
	helper.TestDBSetup(testdb)

	// exec test ~/go/src/gin-gorm-rails-like-sample-api/and remove test db
	code := ExecTest(m, testdb)
	fmt.Println(code)
	fmt.Println(pool)
}
func TestShop(t *testing.T) {
	client := new(http.Client)
	router := server.Router()
	testServer := httptest.NewServer(router)
	defer testServer.Close()

	// GET /shop
	t.Run("returns 200 when GET /", func(t *testing.T) {
		req, _ := http.NewRequest("GET", testServer.URL+"/api/v1/shops", nil)
		res, _ := client.Do(req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error("[Error]", body, err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		expectedBody := `{"shops":[{"id":1,"shop_name":"test shop name","shop_description":"test shop description"}]}`
		assert.JSONEq(t, expectedBody, string(body))
	})

	// GET /shop/:id
	t.Run("returns 200 when GET /1", func(t *testing.T) {
		req, _ := http.NewRequest("GET", testServer.URL+"/api/v1/shops/1", nil)
		res, _ := client.Do(req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error("[Error]", body, err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		expectedBody := `{"id":1,"shop_name":"test shop name","shop_description":"test shop description"}`
		assert.JSONEq(t, expectedBody, string(body))
	})

	t.Run("returns 404 when GET /100 (not existing id)", func(t *testing.T) {
		req, _ := http.NewRequest("GET", testServer.URL+"/api/v1/shops/100", nil)
		res, _ := client.Do(req)
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error("[Error]", body, err)
		}

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}
