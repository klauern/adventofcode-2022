package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}

	fileStr := string(file)
	fmt.Printf("bvwbjplbgvbhsrlpgdmjqwftvncz: %d\n", packetStart("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	fmt.Printf("nppdvjthqldpwncqszvftbrmjlhg: %d\n", packetStart("nppdvjthqldpwncqszvftbrmjlhg", 4))
	fmt.Printf("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg: %d\n", packetStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	fmt.Printf("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw: %d\n", packetStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))

	position := packetStart(fileStr, 4)
	fmt.Println(position)
	fmt.Println(packetStart(fileStr, 14))
}

func packetStart(fileStr string, length int) int {
	for pos, _ := range fileStr[3:] {
		charSet := make(map[byte]bool)
		for i := 0; i < length; i++ {
			if _, ok := charSet[fileStr[(pos+length-1)-i]]; !ok {
				charSet[fileStr[(pos+length-1)-i]] = true
			}
		}
		if len(charSet) == length {
			return pos + length
		}

	}
	return -1
}
