package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	exampleMap := `30373
25512
65332
33549
35390`
	exampleLines := strings.Split(exampleMap, "\n")
	fmt.Printf("Example Map has %d visible\n", totalVisibleTrees(readMap(exampleLines)))

	file, _ := os.ReadFile("day8/input.txt")
	lines := strings.Split(string(file), "\n")
	m := readMap(lines)
	fmt.Println("totalVisibleTrees", totalVisibleTrees(m))

}

func totalVisibleTrees(trees [][]int) int {
	total := 0
	for y, row := range trees {
		for x, height := range row {
			col := zipColumn(trees, x)
			if isUpVisible(col, height, y) {
				total++
				continue
			}
			if isRightVisible(trees[y], height, x) {
				total++
				continue
			}
			if isDownVisible(col, height, y) {
				total++
				continue
			}
			if isLeftVisible(trees[y], height, x) {
				total++
				continue
			}
		}
	}

	return total
}

func zipColumn(trees [][]int, col int) []int {
	res := make([]int, len(trees))
	for i := 0; i < len(trees); i++ {
		res[i] = trees[i][col]
	}
	return res
}

func isDownVisible(col []int, height, y int) bool {
	for i := y + 1; i < len(col); i++ {
		if col[i] >= height {
			return false
		}
	}
	return true
}

func isUpVisible(col []int, height, y int) bool {
	for i := y - 1; i >= 0; i-- {
		if col[i] >= height {
			return false
		}
	}
	return true
}

func isRightVisible(row []int, height, x int) bool {
	for i := x + 1; i < len(row); i++ {
		if row[i] >= height {
			return false
		}
	}
	return true
}

func isLeftVisible(row []int, height, x int) bool {
	for i := x - 1; i >= 0; i-- {
		if row[i] >= height {
			return false
		}
	}
	return true
}

func readMap(lines []string) [][]int {
	m := make([][]int, len(lines))
	for i, line := range lines {
		for _, r := range line {
			val, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			m[i] = append(m[i], val)
		}
	}

	return m
}
