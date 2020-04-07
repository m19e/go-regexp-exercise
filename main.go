package main // import "reg"

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`golang`)
	fmt.Println(r)
}