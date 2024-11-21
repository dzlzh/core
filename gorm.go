package core

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var G_DB *gorm.DB

func NewGorm(driver, dsn string) {
	opens := map[string]func(dsn string) gorm.Dialector{
		"sqlite": sqlite.Open,
		"mysql":  mysql.Open,
	}
	open, ok := opens[strings.ToLower(driver)]
	if !ok {
		panic("not support driver: " + driver)
	}

	var err error

	G_DB, err = gorm.Open(open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
}
