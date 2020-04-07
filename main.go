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

	// 3 repeat "+"
	checkRegexp(`c+x`, "climax")
	checkRegexp(`c+x`, "cx")
	checkRegexp(`c+x`, "cccccx")
	checkRegexp(`c+x`, "x")
	checkRegexp(`c+x`, "climaxxxxx")
	checkRegexp(`c+x`, "hcg")

	// 4 repeat "?"
	checkRegexp(`c?x`, "climax")
	checkRegexp(`c?x`, "cx")
	checkRegexp(`c?x`, "cccccx")
	checkRegexp(`c?x`, "x")
	checkRegexp(`c?x`, "climaxxxxx")
	checkRegexp(`c?x`, "hcg")

	// 5 number, alphabet
	checkRegexp(`[HCG]`, "C")
	checkRegexp(`[HCG]`, "H")
	checkRegexp(`[HCG]`, "A")

		// range
	checkRegexp(`[0-9]`, "5")
	checkRegexp(`[0-9]`, "S")
	checkRegexp(`[A-Z]`, "S")
	checkRegexp(`[A-Z]`, "5")
	checkRegexp(`[A-Z]`, "j")

		// negation
	checkRegexp(`[^0-9]`, "S")
	checkRegexp(`[^0-9]`, "5")

	// 6 replace
	copy := "Boys are always climax after school!"
	r2 := regexp.MustCompile(`[A-Z][a-z]*s`)
	clean := r2.ReplaceAllString(copy, "Girls")

	fmt.Println(clean)
}