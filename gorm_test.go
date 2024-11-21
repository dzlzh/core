package core

import "fmt"

func ExampleNewGorm() {
	NewGorm("sqlite", "test.db")
	fmt.Printf("%T\n", G_DB)
	db, _ := G_DB.DB()
	fmt.Println(db.Ping())
	// Output:
	// *gorm.DB
	// <nil>
}
