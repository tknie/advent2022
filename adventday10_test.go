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

var curCycle = 0
var curIndex = 0
var curCTPos = 0
var screen = make([]byte, 240)
var registerX = 1
var signalStrength = 0
var content = make([]string, 0)

func TestAdventGame10a(t *testing.T) {
	fInput, ferr := os.Open("files/input.10.test2")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()

	for i := 0; i < len(screen); i++ {
		if i < 3 {
			screen[i] = '#'
		} else {
			screen[i] = '.'
		}
	}

	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		content = append(content, value)
		curIndex++
		//fmt.Println(value)
		x := strings.Split(value, " ")
		switch x[0] {
		case "noop":
			curCycle = incCycle(value, curCycle)
		case "addx":
			curCycle = incCycle(value+" pre", curCycle)
			vi, err := strconv.Atoi(x[1])
			if err != nil {
				log.Fatal("addx error value " + x[1])
			}
			curCycle = incCycle(value+" post", curCycle)
			registerX += vi
		}
	}
	dumpScreen()
	fmt.Println("Cycles:", curCycle)
	fmt.Println("Register X:", registerX)
	fmt.Println("Signal strength:", signalStrength)
}

func incCycle(cmd string, curCycle int) int {
	c := curCycle + 1
	if (c-20)%40 == 0 {
		fmt.Println("Cycle:", c, "CT pos", curCTPos, "Cmd:", cmd, "Register X:", registerX, curIndex)
		fmt.Println("Signal strength:", c*registerX)
		signalStrength += (c * registerX)
	}
	curCTPos = curCycle
	p := curCTPos % 40
	// fmt.Println("CT:", cmd, curCTPos, registerX-1, registerX+1, curCTPos >= registerX-1 && curCTPos <= registerX+1)
	if p >= registerX-1 && p <= registerX+1 {
		screen[curCTPos] = '#'
	} else {
		screen[curCTPos] = '.'
	}
	// if (c-20)%40 == 0 || c == 21 {
	// 	fmt.Println("Cycle:", c, "CT pos", curCTPos)
	// 	dumpScreen()
	// }
	return c
}

func dumpScreen() {
	fmt.Println("Dump screen")
	for i, b := range screen {
		fmt.Printf("%c", b)
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func TestAdventGame10b(t *testing.T) {
	fInput, ferr := os.Open("files/input.10.save")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()

	for i := 0; i < len(screen); i++ {
		if i < 3 {
			screen[i] = '#'
		} else {
			screen[i] = '.'
		}
	}

	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		curIndex++
		//fmt.Println(value)
		x := strings.Split(value, " ")
		switch x[0] {
		case "noop":
			curCycle = incCycle(value, curCycle)
		case "addx":
			curCycle = incCycle(value+" pre", curCycle)
			vi, err := strconv.Atoi(x[1])
			if err != nil {
				log.Fatal("addx error value " + x[1])
			}
			curCycle = incCycle(value+" post", curCycle)
			registerX += vi
		}
	}
	fmt.Println("Part One: Signal strength:", signalStrength)
	fmt.Println("Part Two:")
	dumpScreen()
	fmt.Println("Cycles:", curCycle)
	fmt.Println("Register X:", registerX)
}
