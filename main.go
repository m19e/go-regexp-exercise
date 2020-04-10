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

	// 7 numbering match
	colors := "BPRYG"
	r3 := regexp.MustCompile(`BP([A-Z])Y([A-Z])`)
	colors = r3.ReplaceAllString(colors, "BP\x1b[31m${1}\x1b[0mY\x1b[32m${2}\x1b[0m")

	fmt.Println(colors)

	// 8 include "\n" match
	checkRegexp(`^(ko.*)$`, "kaho\nkomiya")
	checkRegexp(`(?m)^(ar.*)$`, "natsuha\narisugawa")

	// 9 split pattern e.g. csv or tsv
	data := "Pink,SONODA Chiyoko, Chiba"
	r4 := regexp.MustCompile(`\s*,\s*`)
	mold := r4.Split(data, -1)

	fmt.Println(mold[0])
	fmt.Println(mold[1])
	fmt.Println(mold[2])

	// 10 match to Slice
	info := "kaho:12-163-45 natsuha:20-168-49 rinze:16-155-44 juri:17-160-48 chiyoko:17-149-46"
	r5 := regexp.MustCompile(`[\d\-]+`)
	fmt.Println(r5.FindAllStringSubmatch(info, -1))

	r6 := regexp.MustCompile(`(\S+):([\d\-]+)`)
	cast := r6.FindAllStringSubmatch(info, -1)
	for _, d := range cast {
		fmt.Println(d[1], d[2])
	}
}
