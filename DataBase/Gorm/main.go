package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  uint
}

func main() {
	// Update the database connection configuration
	dsn := "host=localhost user=postgres password=athun123 dbname=gormtest port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto-migrate the User struct to create the "users" table
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new user
	user := User{Name: "John Doe", Age: 25}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	// Query the user from the database
	var queriedUser User
	result = db.First(&queriedUser)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Queried User:", queriedUser)

	// Update the user's name
	result = db.Model(&queriedUser).Update("Name", "Jane Smith")
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Updated User:", queriedUser)

	// Delete the user from the database
	result = db.Delete(&queriedUser)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	fmt.Println("Deleted User:", queriedUser)
}
