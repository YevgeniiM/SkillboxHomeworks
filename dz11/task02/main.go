package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "a10 10 20b 20 30c30 30 dd"
	var s1, result string
	var spaceIndex int
	fmt.Println("____________________")

	for len(s) > 0 {
		s = strings.Trim(s, " ")
		spaceIndex = strings.Index(s, " ")
		if spaceIndex == -1 {
			s1 = s
			s = ""

		} else {
			s1 = s[:spaceIndex]
			s = s[spaceIndex:]
		}
		if _, err := strconv.Atoi(s1); err == nil {
			result += s1 + " "
		}
	}
	fmt.Println(result)
}
