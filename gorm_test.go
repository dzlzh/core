package core

import "fmt"

func ExampleNewGorm() {
	gorm := NewGorm("sqlite", "test.db")
	fmt.Printf("%T\n", gorm)
	db, _ := gorm.DB()
	fmt.Println(db.Ping())
	// Output:
	// *gorm.DB
	// <nil>
}
