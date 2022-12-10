package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	moves := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	exampleMoves := readVectors(strings.Split(moves, "\n"))
	head, tail := &Knot{
		x:         100,
		y:         100,
		positions: map[string]bool{},
	}, &Knot{
		x: 100,
		y: 100,
		positions: map[string]bool{
			"100,100": true,
		},
	}
	for _, movement := range exampleMoves {
		chase(head, tail, movement)
	}

	fmt.Println("head", head)
	fmt.Println("tail", tail)
	fmt.Println("total positions tail occupied", len(tail.positions))

	file, _ := os.ReadFile("day9/input.txt")
	lines := strings.Split(string(file), "\n")

	movements := readVectors(lines)

	head, tail = &Knot{
		x:         100,
		y:         100,
		positions: map[string]bool{},
	}, &Knot{
		x: 100,
		y: 100,
		positions: map[string]bool{
			"100,100": true,
		},
	}
	for _, movement := range movements {
		chase(head, tail, movement)
	}

	fmt.Println("head", head)
	fmt.Println("tail", tail)
	fmt.Println("total positions tail occupied", len(tail.positions))

	testLarger()
}

func (k *Knot) nearby(from *Knot) bool {
	xDelta := math.Abs(float64(k.x) - float64(from.x))
	yDelta := math.Abs(float64(k.y) - float64(from.y))
	return xDelta <= 1.0 && yDelta <= 1.0
}

func chase(head, tail *Knot, move Movement) {
	for i := 0; i < move.Amount; i++ {
		singleMove(head, tail, move.Direction)
	}
}

func singleMove(head *Knot, tail *Knot, direction string) {
	switch direction {
	case "U":
		head.y += 1
	case "D":
		head.y -= 1
	case "L":
		head.x -= 1
	case "R":
		head.x += 1
	}
	if !tail.nearby(head) {
		// find what move to make U, UL, L, DL, D, DR, R, UR
		xDelta := math.Abs(float64(head.x) - float64(tail.x))
		yDelta := math.Abs(float64(head.y) - float64(tail.y))

		if xDelta >= 2.0 {
			tail.y = head.y
			if tail.x > head.x {
				tail.x = head.x + 1
			} else if tail.x < head.x {
				tail.x = head.x - 1
			}
		}
		if yDelta >= 2.0 {
			tail.x = head.x
			if tail.y > head.y {
				tail.y = head.y + 1
			} else if tail.y < head.y {
				tail.y = head.y - 1
			}
		}
		tail.positions[tail.String()] = true
	}
}

type Vector struct {
	x int
	y int
}

type Movement struct {
	Direction string
	Amount    int
}

type Knot struct {
	x         int
	y         int
	positions map[string]bool
}

func (k *Knot) String() string {
	return fmt.Sprintf("%v,%v", k.x, k.y)
}

func readVectors(lines []string) []Movement {
	var movements []Movement
	for _, line := range lines {
		instruction := strings.Split(line, " ")
		amt, _ := strconv.Atoi(instruction[1])
		movements = append(movements, Movement{
			Direction: instruction[0],
			Amount:    amt,
		})
	}
	return movements
}

func testLarger() {
	moves := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	examples := strings.Split(moves, "\n")
	exampleMoves := readVectors(examples)

	snakes := make([]*Knot, 10)
	for i, _ := range snakes {
		snakes[i] = &Knot{
			x: 0,
			y: 0,
			positions: map[string]bool{
				"0,0": true,
			},
		}
	}

	for _, move := range exampleMoves {
		for i, _ := range snakes {
			if i >= 1 {
				for j := 0; j < move.Amount; j++ {
					for k := i; k > 0; k-- {
						singleMove(snakes[k], snakes[k-1], move.Direction)
					}
				}
			}
		}
	}

	fmt.Println("total positions tail occupied larger example", len(snakes[9].positions))
}
