package main

import "fmt"

func little1(x *int) {
	*x += 10
	fmt.Println("first func")
}
func little2(x *int) {
	*x /= 2
	fmt.Println("second func")
}

func big(x *int, little1 func(*int), little2 func(*int)) {
	little2(x)
	little1(x)

}

func main() {
	a := 2
	x := &a

	little1(x)
	fmt.Println(a)

	little2(x)
	fmt.Println(a)

	big(x, little1, little2)
	fmt.Println(a)
}
