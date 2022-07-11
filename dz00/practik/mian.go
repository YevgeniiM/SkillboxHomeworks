package main

import "fmt"

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB /*ByteSize*/ = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main()  {
	fmt.Println(KB)
	fmt.Println(GB)
	fmt.Println(EB)
}
