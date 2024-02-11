package main

import (
	"fmt"

	"github.com/mohieey/gault"
)

func main() {
	v := gault.Memory("123abc")
	v.Set("secret", "sdsdsds")
	fmt.Println(v.Get("secret"))
}
