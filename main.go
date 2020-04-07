package main // import "reg"

import (
	"fmt"
	"regexp"
)

func checkRegexp(reg, s string) {
	fmt.Printf("「%s」match「%s」? =>", reg, s)
	fmt.Println(regexp.MustCompile(reg).Match([]byte(s)))
}

func main() {
	// 0
	r := regexp.MustCompile(`golang`)
	fmt.Println(r)

	// 1
	r1 := regexp.MustCompile(`world`)
	fmt.Println(r1.MatchString("hello"))
	fmt.Println(r1.MatchString("hello, world"))
}