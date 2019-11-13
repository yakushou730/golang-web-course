package main

import (
	"fmt"

	"github.com/yakushou730/golang-web-course/rand"
)

func main() {
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}
