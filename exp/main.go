package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "TsengYaoShang"
	password = ""
	dbname   = "golang_test"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserId      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var user User
	db.Preload("Orders").First(&user)
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println("Email:", user.Email)
	fmt.Println("Number of orders:", len(user.Orders))
	fmt.Println("Orders:", user.Orders)
}
