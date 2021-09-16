package main

import (
	"fmt"
	"math/rand"
)

func main() {
	stop := 98

	for i := 100; i > 0; i-- {

		if rand.Intn(100) == stop {
			fmt.Println("stop", stop)
			return
		}
		fmt.Println(i)
	}
	fmt.Println("Start")
}
