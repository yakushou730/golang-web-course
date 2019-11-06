package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "TsengYaoShang"
	dbname = "golang_test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// _, err = db.Exec(`
	// INSERT INTO users(name, email)
	// VALUES($1, $2)`,
	// 	"yakushou", "yakushou730@gmail.com")
	// if err != nil {
	// 	panic(err)
	// }

	var id int
	row := db.QueryRow(`
	INSERT INTO users(name, email)
	VALUES($1, $2) RETURNING id`,
		"yakushou", "yakushou730@gmail.com")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id)

	db.Close()
}
