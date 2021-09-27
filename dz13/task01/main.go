package main

import "fmt"

func little1(x *int) {
	*x++
	fmt.Println("first func")
}
func little2(s *string) {
	*s = "Значение переменной равно"
	fmt.Println("second func")
}

func big(little1 func(*int), little2 func(*string)) {
	x := 2
	s := ""
	little2(&s)
	little1(&x)
	fmt.Printf("%s : %d\n", s, x)
}

func main() {
	a := 10
	b := ""

	little1(&a)
	fmt.Println(a)

	little2(&b)
	fmt.Println(b)

	big(little1, little2)
}
