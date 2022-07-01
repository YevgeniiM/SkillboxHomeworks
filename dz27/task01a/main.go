package main

import (
	"fmt"
	"testmod/person"
)

func main() {
	p := person.Person{20,
		"Tom"}
	fmt.Println(p)
}
