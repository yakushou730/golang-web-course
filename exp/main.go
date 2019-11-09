package main

import (
	"fmt"

	"github.com/yakushou730/golang-web-course/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "TsengYaoShang"
	password = ""
	dbname   = "golang_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	// This will error because you DO NOT have a user with
	// this ID, but we will create on soon.
	user, err := us.ByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
