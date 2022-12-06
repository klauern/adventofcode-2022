package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day5/input.txt")
	if err != nil {
		panic(err)
	}

	var stacks []*list.List

	lines := strings.Split(string(file), "\n")
	instructionStart := 0
	for i, line := range lines {
		if line == "" {
			stacks = loadStacks(lines[0 : i-1])
			instructionStart = i
			break
		}
	}

	fmt.Println("instructionStart", lines[instructionStart+1])
	for _, line := range lines[instructionStart+1:] {
		move(stacks, line)
	}

	for _, stack := range stacks {
		fmt.Printf("%s", stack.Front().Value)
	}

}

type Container string

/*
[S]                 [T] [Q]
[L]             [B] [M] [P]     [T]
[F]     [S]     [Z] [N] [S]     [R]
[Z] [R] [N]     [R] [D] [F]     [V]
[D] [Z] [H] [J] [W] [G] [W]     [G]
[B] [M] [C] [F] [H] [Z] [N] [R] [L]
[R] [B] [L] [C] [G] [J] [L] [Z] [C]
[H] [T] [Z] [S] [P] [V] [G] [M] [M]

	1   2   3   4   5   6   7   8   9

1   5   9   13  17  21  25  29  33
*/
func loadStacks(lines []string) []*list.List {
	stacks := make([]*list.List, 9)
	for i := 0; i < 9; i++ {
		stacks[i] = list.New()
	}
	for _, line := range lines {
		for i := 1; i <= 33; i += 4 {
			if line[i] != ' ' {
				position := (i - 1) / 4
				fmt.Printf("Stack at position %d contains %v\n", i, string(line[i]))
				container := Container(line[i])
				stacks[position].PushBack(container)
			}
		}
	}

	return stacks
}

func move(stacks []*list.List, instruction string) {
	// move 6 from 1 to 7
	pieces := strings.Split(instruction, " ")
	amt, from, to := pieces[1], pieces[3], pieces[5]
	amtInt, _ := strconv.Atoi(amt)
	fromInt, _ := strconv.Atoi(from)
	toInt, _ := strconv.Atoi(to)
	for i := 1; i <= amtInt; i++ {
		// move from one stack to another one
		fromStack := stacks[fromInt-1]
		toStack := stacks[toInt-1]
		moveContainer(fromStack, toStack)
	}
}

func moveContainer(from, to *list.List) {
	container := from.Remove(from.Front())
	to.PushFront(container)
}
