package advent

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"testing"
)

type elf struct {
	amount int
	nr     int
}

func TestAdventGame1a(t *testing.T) {
	f, ferr := os.Open("files/input.1")
	if ferr != nil {
		return
	}

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	x := make([]elf, 0)
	current := 0
	higest := 0
	highelf := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		if value == "" {
			e := elf{amount: current, nr: len(x) + 1}
			x = append(x, e)
			if higest < current {
				higest = current
				highelf = len(x)
			}
			current = 0
		} else {
			x, err := strconv.Atoi(value)
			if err != nil {
				return
			}
			current += x
		}
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].amount < x[j].amount
	})

	fmt.Println("Final list:", x)
	fmt.Println("Highest:", higest)
	fmt.Println("Highest elf:", highelf)
	xxx := x[len(x)-1].amount + x[len(x)-2].amount + x[len(x)-3].amount
	fmt.Println("Total", xxx)

	f.Close()

}
