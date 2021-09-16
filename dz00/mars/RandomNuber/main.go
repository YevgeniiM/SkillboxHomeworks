package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	day := time.Now().Weekday()
	fmt.Println(day)
	t := time.Now().Unix()
	fmt.Printf("Текущий timestamp %d\n", t)
	delta := t % 10
	fmt.Println(delta)
	lucky := rand.Intn(int(delta + 5))
	fmt.Println(lucky)
}
