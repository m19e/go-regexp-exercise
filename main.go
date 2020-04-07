package main // import "reg"

import (
	"fmt"
	"regexp"
)

func checkRegexp(reg, s string) {
	fmt.Printf("「%s」match「%s」? => ", reg, s)
	fmt.Println(regexp.MustCompile(reg).Match([]byte(s)))
}

func main() {
	// 0 hello
	r := regexp.MustCompile(`golang`)
	fmt.Println(r)

	// 1 string
	r1 := regexp.MustCompile(`world`)
	fmt.Println(r1.MatchString("hello"))
	fmt.Println(r1.MatchString("hello, world"))

	// 2 repeat "*"
	checkRegexp(`c*x`, "climax")
	checkRegexp(`c*x`, "cx")
	checkRegexp(`c*x`, "cccccx")
	checkRegexp(`c*x`, "x")
	checkRegexp(`c*x`, "climaxxxxx")
	checkRegexp(`c*x`, "hcg")
}