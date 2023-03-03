package main

import (
	"fmt"
)

func hasOverlap(c1 []int, c2 []int, n int) bool {
	for i := 0; i < n; i++ {
		if c1[i] == 1 && c2[i] == 1 {
			return true
		}
	}
	return false
}

func countValidCombinations(winningCombinations [][]int) int {
	const n = 14
	const k = 4

	// Початкова генерація всіх комбінацій
	combinations := make([][]int, 0)
	for i := 0; i < 1<<n; i++ {
		c := make([]int, n)
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				c[j] = 1
			}
		}
		combinations = append(combinations, c)
	}

	// Фільтруємо всі комбінації, щоб виключити ті, які мають 2 або більше спільних цифри з виграшними комбінаціями
	filteredCombinations := make([][]int, 0)
	for _, c := range combinations {
		valid := true
		for _, wc := range winningCombinations {
			if hasOverlap(c, wc, n) && hasOverlap(c, wc, n-2) {
				valid = false
				break
			}
		}
		if valid {
			filteredCombinations = append(filteredCombinations, c)
		}
	}

	// Підрахунок кількості дійсно унікальних комбінацій
	count := 0
	for i := 0; i < len(filteredCombinations); i++ {
		for j := i + 1; j < len(filteredCombinations); j++ {
			if !hasOverlap(filteredCombinations[i], filteredCombinations[j], n) {
				count++
			}
		}
	}

	return count
}

func main() {
	winningCombinations := [][]int{
		{1, 3, 6, 13},
		{2, 6, 11, 7},
		{3, 5, 7, 10},
		{9, 2, 12, 7},
		{10, 9, 7, 5},
		{8, 5, 12, 11},
		{3, 9, 12, 8},
		{9, 6, 3, 14},
		{14, 12, 2, 7},
		{12, 2, 3, 14},
		{9, 5, 6, 4},
	}

	count := countValidCombinations(winningCombinations)
	fmt.Printf("Кількість унікальних комбінацій: %d\n", count)
}

