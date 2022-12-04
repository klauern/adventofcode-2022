package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("day4/input.txt")
	lines := strings.Split(string(file), "\n")
	totalInclusions := 0
	overlaps := 0
	for _, line := range lines {
		first, second := parseLine(line)
		if included(first, second) || included(second, first) {
			fmt.Println(line, "is included")
			totalInclusions++
		}
		if overlap(first, second) || overlap(second, first) {
			overlaps++
		}
	}
	fmt.Println(totalInclusions)
	fmt.Println(overlaps)
}

func included(first *ElvesRange, second *ElvesRange) bool {
	if first.Start >= second.Start && second.End >= first.End {
		if first.End >= second.Start && first.End <= second.End {
			return true
		}
	}
	return false
}

func overlap(first *ElvesRange, second *ElvesRange) bool {
	if first.End >= second.Start && second.End >= first.End {
		return true
	}
	return false
}

type ElvesRange struct {
	Start int
	End   int
}

func parseLine(line string) (*ElvesRange, *ElvesRange) {
	elves := strings.Split(line, ",")
	first := parseRange(elves[0])
	second := parseRange(elves[1])
	return first, second
}

func parseRange(elf string) *ElvesRange {
	ranges := strings.Split(elf, "-")
	start, _ := strconv.Atoi(ranges[0])
	end, _ := strconv.Atoi(ranges[1])
	return &ElvesRange{Start: start, End: end}
}
