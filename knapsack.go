package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"io"
	"strings"
	"strconv"
	"time"
	"sort"
)

type Item struct {
	Value int
	Weight int
}

/**
 * Fully polynomial time approximation scheme.
 */
func KnapsackFPTAS(items []Item, maxWeight int, accuracy float32) ([]int, int) {
	var max int = 0
	var factor float32

	// Find the largest value in all items.
	for i:= 0; i < len(items); i++ {
		if (items[i].Value > max) {
			max = items[i].Value
		}
	}

	if (accuracy >= 1) {
		accuracy = 0
	}

	// Initialize our factor K.
	factor = (1 - accuracy) * (float32(max) / float32(len(items)))

	for i := 0; i < len(items); i++ {
		items[i].Value = int(float32(items[i].Value) / factor)
	}

	configuration, value := KnapsackDynamic(items, maxWeight)

	return configuration, int(float32(value) * factor)
}

/**
 * Dynamic Programming solution.
 */
func KnapsackDynamic(items []Item, maxWeight int) ([]int, int) {
	// Make an array of size n * W.
	var solution = make([][]int, maxWeight + 1)

	for i := 0; i <= maxWeight; i++ {
		solution[i] = make([]int, len(items) + 1)
	}

	// Start our DP solution. O(nW) (Pseudo-polynomial)
	for w := 0; w <= maxWeight; w++ {
		for i := 0; i <= len(items); i++ {
			if (i == 0 || w == 0) {
				solution[w][i] = 0
			} else if (items[i - 1].Weight > w) {
				solution[w][i] = solution[w][i - 1]
			} else {
				solution[w][i] = int(
					math.Max(
						float64(solution[w][i-1]),
						float64(items[i - 1].Value + solution[w - items[i - 1].Weight][i - 1])))
			}
		}
	}

	var configuration []int

	// Backtrack the solution to find the configuration of items taken.
	configuration = BacktrackDynamic(items, maxWeight, solution, maxWeight, len(items), configuration)

	return configuration, solution[maxWeight][len(items)]
}

/**
 * Dynamic Programming solution backtracking helper.
 */
func BacktrackDynamic(items []Item, maxWeight int, solution [][]int, indexWeight int, indexItem int, config []int) []int {
	if indexItem == 0 {
		return config
	}

	if items[indexItem - 1].Weight == 0 || solution[indexWeight][indexItem] != solution[indexWeight][indexItem - 1] {
		config = BacktrackDynamic(items, maxWeight, solution, indexWeight - items[indexItem - 1].Weight, indexItem - 1, config)
		config = append(config, 1)
	} else {
		config = BacktrackDynamic(items, maxWeight, solution, indexWeight, indexItem - 1, config)
		config = append(config, 0)
	}

	return config
}

type CoefSorter []Item
func (a CoefSorter) Len() int { return len(a) }
func (a CoefSorter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a CoefSorter) Less(i ,j int) bool { return float32(a[i].Value) / float32(a[i].Weight) > float32(a[j].Value) / float32(a[j].Weight) }

func KnapsackHeuristic(items []Item, maxWeight int, sorter sort.Interface) ([]int, int) {
	sort.Sort(sorter)

	var sumWeight int = 0
	var sumValue int = 0
	var bestConfig []int

	for i := 0; i < len(items); i++ {
		currentWeight := items[i].Weight
		if (sumWeight + currentWeight < maxWeight) {
			bestConfig = append(bestConfig, i)
			sumWeight += items[i].Weight
			sumValue += items[i].Value
		}
	}

	return bestConfig, sumValue
}

func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		err error = nil
		line, ln []byte)

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

func main() {
	/** CONFIGURATION */
	var maxWeight int = 99457
	var file string = "data/hard2"
	/** END CONFIGURATION */

	var start time.Time
	var elapsed time.Duration

	// Open the file.
	f, err := os.Open(file)

	if (err != nil) {
		fmt.Println(err)
	}

	r := bufio.NewReader(f)
	s, e := Readln(r)

	fmt.Println(s)

	// Define where we are to be putting all of our items.
	var items []Item

	for e == nil {
		s, e := Readln(r)
		if (e == io.EOF) {
			break
		}
		// Get our Weight/Value pair
		tuple := strings.Fields(s)

		w, e := strconv.ParseInt(tuple[0], 0, 64)
		v, e := strconv.ParseInt(tuple[1], 0, 64)

		items = append(items, Item{Value:int(v), Weight:int(w)})
	}

	fmt.Println("----------------")
	fmt.Println("--- Dynamic ----")
	fmt.Println("----------------")

	fmt.Println("Items in the set: ",  len(items))
	fmt.Println("Max Weight:", maxWeight)

	start = time.Now()
	_, solution  := KnapsackDynamic(items, maxWeight)
	elapsed = time.Since(start)
	fmt.Println("----------------")
	fmt.Printf("Dynamic took %s\n", elapsed)
	// fmt.Println("Items to take: ")
	// fmt.Printf("%v\n", config)

	fmt.Println("Max Value: ", solution)

	// fmt.Println("----------------")
	// fmt.Println("---- FPTAS -----")
	// fmt.Println("----------------")


	// fmt.Println("Items in the set: ",  len(items))
	// fmt.Println("Max Weight: ", maxWeight)

	// start = time.Now()
	// _, solution = KnapsackFPTAS(items, maxWeight, 0.2)
	// elapsed = time.Since(start)
	// fmt.Println("----------------")
	// fmt.Printf("FPTAS (0.2) took %s\n", elapsed)
	// fmt.Println("Max Value: ", solution)

	// start = time.Now()
	// _, solution = KnapsackFPTAS(items, maxWeight, 0.4)
	// elapsed = time.Since(start)
	// fmt.Println("----------------")
	// fmt.Printf("FPTAS (0.4) took %s\n", elapsed)
	// fmt.Println("Max Value: ", solution)

	// start = time.Now()
	// _, solution = KnapsackFPTAS(items, maxWeight, 0.6)
	// elapsed = time.Since(start)
	// fmt.Println("----------------")
	// fmt.Printf("FPTAS (0.6) took %s\n", elapsed)
	// fmt.Println("Max Value: ", solution)

	// start = time.Now()
	// _, solution = KnapsackFPTAS(items, maxWeight, 0.8)
	// elapsed = time.Since(start)
	// fmt.Println("----------------")
	// fmt.Printf("FPTAS (0.8) took %s\n", elapsed)
	// fmt.Println("Max Value: ", solution)

	fmt.Println("----------------")
	fmt.Println("--- Heuristic --")
	fmt.Println("----------------")

	start = time.Now()
	_, solution = KnapsackHeuristic(items, maxWeight, CoefSorter(items))
	elapsed = time.Since(start)
	fmt.Println("----------------")
	fmt.Printf("Heuristic search took %s\n", elapsed)
	fmt.Println("Max Value: ", solution)
}