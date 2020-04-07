package main // import "reg"

import (
	"fmt"
	"regexp"
)

func main() {
	// 0
	r := regexp.MustCompile(`golang`)
	fmt.Println(r)

	// 1
	r1 := regexp.MustCompile(`world`)
	fmt.Println(r1.MatchString("hello"))
	fmt.Println(r1.MatchString("hello, world"))
}