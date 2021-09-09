package main

import (
	"fmt"
	"math"
)

func main(){
	var i, i8, i16  int
	for i=1; i <= math.MaxUint32; i++ {
		if i % math.MaxUint8 == 0 {
			i8 ++
		}
		if i % math.MaxUint16 == 0 {
			i16 ++
		}
	}
	fmt.Println(math.MaxUint32 / math.MaxUint8)
	fmt.Println("в диапазоне от 0 до uint32 приходится переполнений чисел типа uint8",i8,"и переполнений чисел типа uint16", i16)
}
