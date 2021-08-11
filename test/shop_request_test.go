package test

import (
	"fmt"
	"gin-gorm-rails-like-sample-api/db"
	"gin-gorm-rails-like-sample-api/server"
	"gin-gorm-rails-like-sample-api/test/helper"
	test_db "gin-gorm-rails-like-sample-api/test/testdata/db"
	"io/ioutil"
	"log"
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

	t.Run("returns 200 when GET /", func(t *testing.T) {
		req, _ := http.NewRequest("GET", testServer.URL+"/api/v1/shops", nil)
		res, _ := client.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		log.Printf("ResponseBody:%s\n", string(body))

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

}
