package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type pairs struct {
	x        [100]byte
	beg, end int
}

func TestAdventGame4a(t *testing.T) {
	fInput, ferr := os.Open("files/input.4.save")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	containsPairs := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		first := strings.Split(value, ",")
		x := extractRange(first[0])
		y := extractRange(first[1])
		// fmt.Println(x)
		// fmt.Println(y)
		if x.beg <= y.beg && x.end >= y.end {
			containsPairs++
			fmt.Println("1+")
		} else if y.beg <= x.beg && y.end >= x.end {
			containsPairs++
			fmt.Println("1*")
		} else {
			fmt.Println("0.")
		}
	}
	fmt.Println("Contains pairs:", containsPairs)
}

func TestAdventGame4b(t *testing.T) {
	fInput, ferr := os.Open("files/input.4.save")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	containsPairs := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		first := strings.Split(value, ",")
		x := extractRange(first[0])
		y := extractRange(first[1])
		// fmt.Println(x)
		// fmt.Println(y)
		if x.beg <= y.beg && x.end >= y.end {
			containsPairs++
			fmt.Println("1+")
		} else if y.beg <= x.beg && y.end >= x.end {
			containsPairs++
			fmt.Println("1*")
		} else if y.end < x.beg || x.end < y.beg {
			fmt.Println("0.")
		} else {
			containsPairs++
			fmt.Println("1x")
		}
	}
	fmt.Println("Contains pairs:", containsPairs)
}

func extractRange(value string) pairs {
	s := strings.Split(value, "-")
	beg, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal("Errr " + err.Error())
	}
	end, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal("Errr " + err.Error())
	}
	var x [100]byte
	if beg > 100 || end > 100 {
		log.Fatal("Execeed", beg, end)
	}
	for i := 0; i < 100; i++ {
		if beg <= i+1 && i+1 <= end {
			x[i] = byte(i + 1)
		} else {
			x[i] = 0
		}
	}
	return pairs{x, beg, end}
}
