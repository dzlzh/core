package core

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGorm(driver, dns string) *gorm.DB {
	opens := map[string]func(dsn string) gorm.Dialector{
		"sqlite": sqlite.Open,
		"mysql":  mysql.Open,
	}
	open, ok := opens[strings.ToLower(driver)]
	if !ok {
		panic("not support driver: " + driver)
	}

	db, err := gorm.Open(open(dns), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	return db
}
