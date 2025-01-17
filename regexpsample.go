package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" // string to search
	pat := "[0-9]+.[0-9]+"                                      // pattern search in searchIn

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		fmt.Println("v", v)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}
	re, _ := regexp.Compile(pat)

	str := re.ReplaceAllString(searchIn, "##.#") // replace pat with "##.#"
	fmt.Println(str)
	// using a function
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
