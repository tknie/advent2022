package advent

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

type game struct {
	card int
	game int
	x    string
}

func TestAdventGame2a(t *testing.T) {
	f, ferr := os.Open("files/input.2")
	if ferr != nil {
		return
	}

	rounds := make([]game, 0)
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	round := 0
	total := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		if value == "" {
			continue
		}
		g := game{x: value}
		// r := round % 3
		// switch r {
		// case 0:
		// 	g.game = 6
		// case 1:
		// 	g.game = 0
		// case 2:
		// 	g.game = 3

		// }
		x := strings.Split(value, " ")
		if len(x) > 1 {
			fmt.Println(x[1])
			switch strings.ToUpper(x[1]) {
			case "A", "X": // Stein
				if x[0] == "A" {
					g.game = 3
				}
				if x[0] == "B" {
					g.game = 0
				}
				if x[0] == "C" {
					g.game = 6
				}

				g.card = 1
			case "B", "Y": // Papier
				if x[0] == "A" {
					g.game = 6
				}
				if x[0] == "B" {
					g.game = 3
				}
				if x[0] == "C" {
					g.game = 0
				}
				g.card = 2
			case "C", "Z": // Schere
				if x[0] == "A" {
					g.game = 0
				}
				if x[0] == "B" {
					g.game = 6
				}
				if x[0] == "C" {
					g.game = 3
				}
				g.card = 3
			}
			total += g.card + g.game
		}
		rounds = append(rounds, g)
		round++
	}
	for _, g := range rounds {
		fmt.Println(g)
	}
	fmt.Println("Total:", total)
	f.Close()
}

func TestAdventGame2b(t *testing.T) {
	f, ferr := os.Open("files/input.2")
	if ferr != nil {
		return
	}

	rounds := make([]game, 0)
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	round := 0
	total := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		if value == "" {
			continue
		}
		g := game{x: value}
		// r := round % 3
		// switch r {
		// case 0:
		// 	g.game = 6
		// case 1:
		// 	g.game = 0
		// case 2:
		// 	g.game = 3

		// }
		x := strings.Split(value, " ")
		if len(x) > 1 {
			fmt.Println(x[1])
			switch strings.ToUpper(x[0]) {
			case "A", "X": // Stein
				if x[1] == "X" {
					g.game = 0
					g.card = 3 // Scissors
				}
				if x[1] == "Y" {
					g.game = 3
					g.card = 1 // Rock
				}
				if x[1] == "Z" {
					g.game = 6
					g.card = 2 // Paper
				}

			case "B", "Y": // Papier
				if x[1] == "X" {
					g.game = 0
					g.card = 1
				}
				if x[1] == "Y" {
					g.game = 3
					g.card = 2
				}
				if x[1] == "Z" {
					g.game = 6
					g.card = 3
				}
			case "C", "Z": // Schere
				if x[1] == "X" {
					g.game = 0
					g.card = 2
				}
				if x[1] == "Y" {
					g.game = 3
					g.card = 3
				}
				if x[1] == "Z" {
					g.game = 6
					g.card = 1
				}
			}
			total += g.card + g.game
		}
		rounds = append(rounds, g)
		round++
	}
	for _, g := range rounds {
		fmt.Println(g)
	}
	fmt.Println("Total:", total)
	f.Close()
}
