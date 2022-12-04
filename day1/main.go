package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumInputs(lines []string) []int {
	total := 0
	totals := []int{}

	for _, line := range lines {
		if line == "" {
			totals = append(totals, total)
			total = 0
		} else {
			amt, _ := strconv.Atoi(line)
			total += amt
		}
	}
	return totals
}

func main() {
	file, err := os.ReadFile("day1/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	totals := sumInputs(lines)
	var max int
	for _, total := range totals {
		if max < total {
			max = total
		}
	}

	sort.Ints(totals)
	fmt.Printf("%d\n", max)
	fmt.Printf("%d %d %d\n", totals[len(totals)-1], totals[len(totals)-2], totals[len(totals)-3])
	fmt.Printf("top 3 total %d\n", totals[len(totals)-3]+totals[len(totals)-2]+totals[len(totals)-1])
}
