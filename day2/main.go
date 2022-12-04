package main

import (
	"fmt"
	"os"
	"strings"
)

type Game struct {
	You   string
	Them  string
	Score int
}

var Points map[string]int = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func NewGame(you, them string) *Game {
	switch Points[you] - Points[them] {
	case -2: // Rock(1) vs Scissors(3)
		return &Game{
			You:   you,
			Them:  them,
			Score: Points[you] + 6, // 6 for the win
		}
	case -1: // Rock(1) vs Paper(2) or Paper(2) vs Scissors(3)
		return &Game{
			You:   you,
			Them:  them,
			Score: Points[you],
		}
	case 0: // Rock(1) vs Rock(1), Paper(2) vs Paper(2), or Scissors(3) vs Scissors(3)
		{
			return &Game{
				You:   you,
				Them:  them,
				Score: Points[you] + 3,
			}
		}
	case 1: // Paper(2) vs Rock(1) or Scissors(3) vs Paper(2)
		return &Game{
			You:   you,
			Them:  them,
			Score: Points[you] + 6, // 6 for the win
		}
	case 2: // Scissors(3) vs Rock(1)
		return &Game{
			You:   you,
			Them:  them,
			Score: Points[you],
		}
	}
	return nil
}

func ExpectedGame(you, them string) *Game {
	switch you {
	case "X": // you need to lose
		switch them {
		case "A": // They chose rock
			return &Game{
				You:   "Z",
				Them:  "A",
				Score: Points["Z"],
			}
		case "B": // they chose paper
			return &Game{
				You:   "X",
				Them:  "B",
				Score: Points["X"],
			}
		case "C": // they chose scissors
			return &Game{
				You:   "Y",
				Them:  "C",
				Score: Points["Y"],
			}
		}
	case "Y": // you need to draw
		switch them {
		case "A":
			return &Game{
				You:   you,
				Them:  "A",
				Score: Points[them] + 3,
			}
		case "B":
			return &Game{
				You:   you,
				Them:  "B",
				Score: Points[them] + 3,
			}
		case "C":
			return &Game{
				You:   you,
				Them:  "C",
				Score: Points[them] + 3,
			}
		}
	case "Z": // you need to win
		switch them {
		case "A": // they chose rock
			return &Game{
				You:   "Y",
				Them:  "A",
				Score: Points["Y"] + 6,
			}
		case "B": // they chose paper
			return &Game{
				You:   "Z",
				Them:  "B",
				Score: Points["Z"] + 6,
			}
		case "C": // they chose scissors
			return &Game{
				You:   "X",
				Them:  "C",
				Score: Points["X"] + 6,
			}
		}
	}
	return nil
}

func main() {
	file, err := os.ReadFile("day2/input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	total := 0
	strategyTotal := 0
	for _, line := range lines {
		players := strings.Split(line, " ")
		you := players[1]
		them := players[0]
		game := NewGame(you, them)
		expected := ExpectedGame(you, them)
		total += game.Score
		strategyTotal += expected.Score
	}
	fmt.Printf("total: %d\n", total)
	fmt.Printf("strategy total: %d\n", strategyTotal)
}
