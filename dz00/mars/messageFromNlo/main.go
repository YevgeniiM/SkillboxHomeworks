package main

import (
	"fmt"
)

func main() {

	//message := "uv vagreangvbany fcnpr fgngvba"

	//for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII
	//	c := message[i]
	//	if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы
	//		c = c + 13
	//		if c > 'z' {
	//			c = c - 26
	//		}
	//	}
	//	fmt.Printf("%c", c)

	//message := "L fdph, L vdz, L frqtxhuhg."
	//
	//	for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII
	//		c := message[i]
	//		if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы
	//			c = c - 3
	//			if c < 'a' {
	//				c = c + 26
	//			}
	//		}
	//		if c >= 'A' && c <= 'Z' { // Оставляет оригинальную пунктуацию и пробелы
	//			c = c - 3
	//			if c < 'A' {
	//				c = c + 26
	//			}
	//		}
	//		fmt.Printf("%c", c)
	//}
	message := "Hola Estación Espacial Internacional"

	for _, c := range message {
		if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		if c >= 'A' && c <= 'Z' { // Оставляет оригинальную пунктуацию и пробелы
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c) // "Hola Estación Espacial Internacional" ROOT13 Ubyn Rfgnpvóa Rfcnpvny Vagreanpvbany

	}
	fmt.Println("")
}
