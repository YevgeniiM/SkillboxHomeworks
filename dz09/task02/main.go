package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int16
	var multiplicationResult int32
	fmt.Println("")
	_, _ = fmt.Scanf("%d,%d", &a, &b)
	multiplicationResult = int32(a) * int32(b)
	fmt.Println("", multiplicationResult)

	switch {
	case multiplicationResult <= math.MaxUint8 && multiplicationResult >= 0:
		fmt.Printf("type %T result %v ", uint8(multiplicationResult), uint8(multiplicationResult))
	case multiplicationResult <= math.MaxUint16 && multiplicationResult >= 0:
		fmt.Printf("type %T result %v ", uint16(multiplicationResult), uint16(multiplicationResult))
	case multiplicationResult >= math.MinInt8 && multiplicationResult <= 0:
		fmt.Printf("type %T result %v ", int8(multiplicationResult), int8(multiplicationResult))
	case multiplicationResult >= math.MinInt16 && multiplicationResult <= 0:
		fmt.Printf("type %T result %v ", int16(multiplicationResult), int16(multiplicationResult))
	default:
		fmt.Printf("type %T result %v ", multiplicationResult, multiplicationResult)
	}
}
