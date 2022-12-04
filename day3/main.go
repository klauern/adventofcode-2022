package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("day3/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	totalPriorities := int32(0)
	// rune is an alias for int32
	for _, line := range lines {
		item, priority := commonItem(line)
		totalPriorities += priority
		fmt.Println(string(item), priority)
	}

	totalGroupPriorities := int32(0)
	for i := 0; i < len(lines); i += 3 {
		_, priority := commonGroup(lines[i], lines[i+1], lines[i+2])
		totalGroupPriorities += priority
	}

	fmt.Println(totalPriorities)
	fmt.Println(totalGroupPriorities)

}

func commonItem(rucksack string) (rune, int32) {
	length := len(rucksack)

	// first compartment
	first := map[rune]bool{}
	for _, r := range rucksack[0 : length/2] {
		first[r] = true
	}

	// second compartment
	second := map[rune]bool{}
	for _, r := range rucksack[length/2:] {
		second[r] = true
	}

	for r, _ := range first {
		if second[r] {
			return r, lookupPriority(r)
		}
	}
	return 'A', 0
}

func commonGroup(first, second, third string) (rune, int32) {
	// first compartment
	firstCompartment := map[rune]bool{}
	for _, r := range first {
		firstCompartment[r] = true
	}

	// second compartment
	secondCompartment := map[rune]bool{}
	for _, r := range second {
		secondCompartment[r] = true
	}

	for _, r := range third {
		if firstCompartment[r] && secondCompartment[r] {
			return r, lookupPriority(r)
		}
	}
	return 'A', 0
}

func lookupPriority(item rune) int32 {
	if unicode.IsUpper(item) {
		return int32(item) - 38
	}
	return int32(item) - 96

}
