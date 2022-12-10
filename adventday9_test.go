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

type point struct {
	x, y int
}

const maxDimension = 1000
const countNodes = 9

func TestAdventGame9a(t *testing.T) {
	fInput, ferr := os.Open("files/input.9.test")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
	}
	finalWord := ""
	fmt.Println("Word:", finalWord)
}

func TestAdventGame9b(t *testing.T) {
	fInput, ferr := os.Open("files/input.9.save")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	matrix := make([][]byte, maxDimension)
	for i := 0; i < maxDimension; i++ {
		matrix[i] = make([]byte, maxDimension)
	}
	// org := &point{500, 500}
	s := &point{maxDimension / 2, maxDimension / 2}
	T := make([]*point, countNodes)
	for i := 0; i < countNodes; i++ {
		T[i] = &point{maxDimension / 2, maxDimension / 2}
	}
	H := &point{maxDimension / 2, maxDimension / 2}
	matrix[T[0].x][T[0].y] = 1

	if maxDimension < 100 {
		dumpMatrix(matrix, s, H, T)
	}

	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		move := strings.Split(value, " ")
		moveLen, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatal("Move len errro " + move[1])
		}
		matrix = matrixLayer(matrix, H, T, move[0], moveLen)
		if maxDimension < 100 {
			dumpMatrix(matrix, s, H, T)
		}
	}
	fmt.Println("s:", s)
	fmt.Println("H:", H)
	fmt.Println("T:", T)
	dumpMatrixUsage(matrix)
}

func moveRope(matrix [][]byte, H *point, T []*point) {
	cH := H
	for _, t := range T {
		moveTail(cH, t)
		cH = t
	}
	matrix[T[len(T)-1].x][T[len(T)-1].y] = 1

}

func matrixLayer(matrix [][]byte, H *point, T []*point, move string, moveLen int) [][]byte {
	fmt.Println("Move", move, moveLen)
	switch move {
	case "R":
		for x := 0; x < moveLen; x++ {
			H.x++
			moveRope(matrix, H, T)
			if maxDimension < 100 {
				dumpMatrix(matrix, H, H, T)
			}
		}
	case "L":
		for x := 0; x < moveLen; x++ {
			H.x--
			moveRope(matrix, H, T)
			if maxDimension < 100 {
				dumpMatrix(matrix, H, H, T)
			}
		}
	case "U":
		for x := 0; x < moveLen; x++ {
			H.y++
			moveRope(matrix, H, T)
			if maxDimension < 100 {
				dumpMatrix(matrix, H, H, T)
			}
		}
	case "D":
		for x := 0; x < moveLen; x++ {
			H.y--
			moveRope(matrix, H, T)
			if maxDimension < 100 {
				dumpMatrix(matrix, H, H, T)
			}
		}
	}
	return matrix
}

func dumpMatrix(matrix [][]byte, s, H *point, T []*point) {
	cH := H
	for y, m := range matrix {
		for x := range m {
			var fT *point
			index := -1
			for i, cT := range T {
				if cT.x == x && cT.y == y {
					fT = cT
					index = i + 1
					break
				}
			}
			switch {
			case cH.x == x && cH.y == y:
				fmt.Print("H")
			case fT != nil:
				fmt.Print(index)
			case s.x == x && s.y == y:
				fmt.Print("s")
			case matrix[x][y] == 1:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func dumpMatrixUsage(matrix [][]byte) {
	usage := 0
	for y, m := range matrix {
		for x := range m {
			switch {
			case matrix[x][y] == 1:
				fmt.Print("#")
				usage++
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("Usage:", usage)
}

func moveTail(head, tail *point) *point {
	dx := head.x - tail.x
	dy := head.y - tail.y
	// horiz
	if dy == 0 {
		if dx == 2 {
			tail.x = tail.x + 1
			//tail.y = tail.y
			return tail
		} else if dx == -2 {
			tail.x = tail.x - 1
			//tail.y = tail.y
			return tail
		}
		return tail
	}
	// vert
	if dx == 0 {
		if dy == 2 {
			//tail.x = tail.x
			tail.y = tail.y + 1
			return tail
		} else if dy == -2 {
			//tail.x = tail.x
			tail.y = tail.y - 1
			return tail
		}
		return tail
	}

	if dx == 2 || dy == 2 || dx == -2 || dy == -2 {
		if dx > 0 {
			tail.x = tail.x + 1
			// tail.y = tail.y
			// tail = tail
		} else {
			tail.x = tail.x - 1
			// tail.y = tail.y
			//tail = tail.West()
		}
		if dy > 0 {
			//tail.x = tail.x
			tail.y = tail.y + 1
			//tail = tail
		} else {
			//tail.x = tail.x
			tail.y = tail.y - 1
			//tail = tail
		}
		return tail
	}
	return tail
}
