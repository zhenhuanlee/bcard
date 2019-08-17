package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// pg
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Setup setup
func Setup() {
	// db, err := sql.Open("postgres", "postgres://jude:@localhost/doge_production")

	// if err != nil {
	// 	panic("init postgresql connection failed")
	// }

	// db.SetMaxOpenConns(10)
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=jude dbname=doge_development sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("init db failed: %s", err))
	}

	db.Select("")
}
