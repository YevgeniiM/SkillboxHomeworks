package main

import "fmt"

//const size = 10

func sortInsert(input []uint) []uint {
	//change := false
	sortM := make([]uint, 0)
	s := make([]uint, len(input))
	for i := 0; i < len(input); i++ {
		for j := 0; j <= i; j++ {
			if j == i {
				sortM = append(sortM, input[i])
			} else if input[i] <= sortM[j] {
				copy(s, sortM[j:i])
				sortM = append(sortM[:j], input[i])
				sortM = append(sortM[:j+1], s[:i-j]...)
				break
			}
		}

	}

	return sortM
}
func main() {

	a := []uint{1, 20, 5, 4, 65, 5, 7, 4, 54, 9}
	m := sortInsert(a)
	fmt.Println(m)

}
