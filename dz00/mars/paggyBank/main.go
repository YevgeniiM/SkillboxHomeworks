package main

import (
	"fmt"
	"math/rand"
)

func main() {
	paggyBank := 0
	for paggyBank < 2500 {
		switch rand.Intn(3) {
		case 0:
			paggyBank += 5
		case 1:
			paggyBank += 10
		case 2:
			paggyBank += 25
		}
		fmt.Printf("In paggyBank %d.%02d$\n", paggyBank/100, paggyBank%100)
		// підрахунок иде в центах, тому долари це  paggyBank/100, а paggyBank%100 цетни і просто виводимо два инта через крапку "." (%02d - заповнює нулі при виводі)
	}

}
