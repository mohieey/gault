package main

import (
	"fmt"

	"github.com/mohieey/gault"
)

func main() {
	v := gault.File("123abc", "sr")
	v.Set("secret", "sdsdsds")
	val, _ := v.Get("secret")
	fmt.Println(val)
}
