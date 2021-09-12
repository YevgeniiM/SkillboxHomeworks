package main

import (
	"fmt"
	"strings"
)

func main() {

	var numb, indexA int
	s := "Golang is programming language"
	for {
		indexA = strings.Index(s, "a")
		//fmt.Println(numb1)
		if indexA >= 0 {
			s = s[indexA+1:]
			numb++
		} else {
			break
		}
	}
	fmt.Println("Количество букв 'a' в строке равно", numb)
}
