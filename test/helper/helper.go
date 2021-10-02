package helper

import (
	"log"

	"gin-gorm-rails-like-sample-api/model/entity"

	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
)

func TestDBSetup(testdb *gorm.DB) {

	testdb.AutoMigrate(&entity.Shop{})

	// Loading fixtures
	db, _ := testdb.DB()

	fixtures, err := testfixtures.New(
		testfixtures.Template(),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("testdata/fixtures"),
	)

	if err != nil {
		log.Fatalf("Fixtures init error: %s", err)
	}
	if err := fixtures.Load(); err != nil {
		log.Fatalf("Fixtures load error: %s", err)
	}
}
