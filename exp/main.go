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

	user := models.User{
		Name:  "yakushou",
		Email: "yakushou730@gmail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	// NOTE: You may need to update the query code a bit as well
	foundUser, err := us.ByEmail("yakushou730@gmail.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(foundUser)
}
