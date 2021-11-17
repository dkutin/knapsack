package main

import (
	"fmt"
	"math"
)

type Item struct {
	Value int
	Weight int
}

/**
 * Fully polynomial time approximation scheme.
 */
func KnapsackFPTAS(items []Item, maxWeight int, accuracy float32) ([]Item, int) {
	var max int = 0
	var factor float32

	// Find the largest value in all items.
	for i:= 0; i < len(items); i++ {
		if (items[i].Value > max) {
			max = items[i].Value
		}
	}

	// Initialize our factor K.
	factor = (1 - factor) * (float32(max) / float32(len(items)))

	for i := 0; i < len(items); i++ {
		items[i].Value = int(float32(items[i].Value) / factor)
	}

	configuration, value := KnapsackDynamic(items, maxWeight)

	return configuration, int(float32(value) * factor)
}

/**
 * Dynamic Programming solution.
 */
func KnapsackDynamic(items []Item, maxWeight int) ([]Item, int) {
	// Make an array of size n * W.
	var solution [][] int

	// Initialize all entries to 0 along the items column. O(n)
	for i := 0; i < maxWeight; i++ {
		solution[0][i] = 0
	}
	// Initialize all entries to 0 along the weight row. O(n)
	for i := 0; i < len(items); i++ {
		solution[i][0] = 0
	}

	// Start our DP solution. O(nW) (Pseudo-polynomial)
	for i := 0; i < len(items); i++ {
		for j := 0; j < maxWeight; j++ {
			if (items[i].Weight > j) {
				solution[i][j] = solution[i-1][j]
			} else {
				solution[i][j] = int(math.Max(float64(solution[i-1][j]), float64(solution[i-1][(j - items[i].Weight)] + items[i].Value)))
			}
		}
	}

	var configuration []Item
	// TODO: We now have some solution. Backtrack to build it.

	return configuration, 1
}

func main() {
	fmt.Println("Lets get started...")
}