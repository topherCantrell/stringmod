package main

import (
	"fmt"

	qt "github.com/topherCantrell/stringmod/quotes"
	str "github.com/topherCantrell/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(qt.GetEmoji("turtle"))
}
