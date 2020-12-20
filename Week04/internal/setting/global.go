package setting

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

var (
	dataConfig = Config{}
	DB         *sql.DB
	once       = new(sync.Once)
)

func init() {
	once.Do(func() {
		connStr := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			dataConfig.User,
			dataConfig.Password,
			dataConfig.Host,
			dataConfig.Name)
		var err error
		DB, err = sql.Open("mysql", connStr)
		if err != nil {
			log.Fatalf("failed connectiong to database: %v", err)
		}
	})
}
