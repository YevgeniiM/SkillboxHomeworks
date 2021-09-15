package main

import (
	"fmt"
	"strings"
)

func main() {
	var endWord1, endWord2, minEndWord, numb int
	var s, s1, s2 string
	s = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
	//s = "    "
	s = strings.Trim(s, " ,")
	fmt.Println(s)

	for len(s) > 0 {
		s = strings.Trim(s, " ,")
		endWord1 = strings.Index(s, " ")
		endWord2 = strings.Index(s, ",")
		minEndWord = endWord1
		if endWord1 == -1 || endWord2 < endWord1 && endWord2 >= 0 {
			minEndWord = endWord2
		}

		if minEndWord == -1 {
			s1 = s
			s = ""
		} else {
			s1 = s[:minEndWord]
			s = s[minEndWord+1:]
		}
		s2 = strings.Title(s1) //s1 = string([]rune(s)[0])
		if s1 == s2 {
			numb++
		}

	}

	fmt.Printf("В тексте находится %v слов начинающихся с заглавных букв", numb)
	//fmt.Println()
}
