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
	exTrees, exScore := totalVisibleTrees(readMap(exampleLines))
	fmt.Printf("Example Map has %d visible with a max scenic score of %d\n", exTrees, exScore)

	file, _ := os.ReadFile("day8/input.txt")
	lines := strings.Split(string(file), "\n")
	m := readMap(lines)
	trees, score := totalVisibleTrees(m)
	fmt.Println("totalVisibleTrees", trees, "score", score)

}

func totalVisibleTrees(trees [][]int) (int, int) {
	total := 0
	maxScore := 0
	for y, row := range trees {
		for x, height := range row {
			col := zipColumn(trees, x)

			up, upScore := isUpVisible(col, height, y)
			right, rightScore := isRightVisible(trees[y], height, x)
			down, downScore := isDownVisible(col, height, y)
			left, leftScore := isLeftVisible(trees[y], height, x)
			if up || left || right || down {
				score := scenicScore(upScore, downScore, leftScore, rightScore)
				fmt.Println("score:", score, upScore, downScore, leftScore, rightScore)
				if score > maxScore {
					maxScore = score
				}
				total += 1
			}
		}
	}

	return total, maxScore
}

func scenicScore(up, down, left, right int) int {
	return up * down * left * right
}

func zipColumn(trees [][]int, col int) []int {
	res := make([]int, len(trees))
	for i := 0; i < len(trees); i++ {
		res[i] = trees[i][col]
	}
	return res
}

func isDownVisible(col []int, height, y int) (bool, int) {
	visibleTrees := 0
	for i := y + 1; i < len(col); i++ {
		if col[i] >= height {
			return false, visibleTrees + 1
		}
		visibleTrees += 1
	}
	return true, visibleTrees
}

func isUpVisible(col []int, height, y int) (bool, int) {
	visibleTrees := 0
	for i := y - 1; i >= 0; i-- {
		if col[i] >= height {
			return false, visibleTrees + 1
		}
		visibleTrees += 1
	}
	return true, visibleTrees
}

func isRightVisible(row []int, height, x int) (bool, int) {
	visibleTrees := 0
	for i := x + 1; i < len(row); i++ {
		if row[i] >= height {
			return false, visibleTrees + 1
		}
		visibleTrees += 1
	}
	return true, visibleTrees
}

func isLeftVisible(row []int, height, x int) (bool, int) {
	visibleTrees := 0
	for i := x - 1; i >= 0; i-- {
		if row[i] >= height {
			return false, visibleTrees + 1
		}
		visibleTrees += 1
	}
	return true, visibleTrees
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
