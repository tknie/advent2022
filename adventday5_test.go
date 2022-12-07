package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type crates struct {
	index int
	v     string
}

type stack struct {
	index int
	st    *Stack
}

type move struct {
	count, from, to int
}

func TestAdventGame5a(t *testing.T) {
	fInput, ferr := os.Open("files/input.5.save")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	moves := false
	stackIndex := make([]*stack, 0)
	moveSteps := make([]move, 0)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)

		if value == "" {
			moves = true
			sort.Slice(stackIndex, func(i, j int) bool {
				return stackIndex[i].index < stackIndex[j].index
			})
			for _, s := range stackIndex {
				fmt.Println(s.index, s.st.Size)
			}
			continue
		}
		if moves {
			from := strings.Index(value, "from")
			to := strings.Index(value, "to")
			sCount, err := strconv.Atoi(value[5 : from-1])
			if err != nil {
				log.Fatal("Error count >" + value[5:from] + "<" + err.Error())
			}
			sNrFrom, err := strconv.Atoi(value[from+5 : to-1])
			if err != nil {
				log.Fatal("Error from >" + value[from+5:to] + "<" + err.Error())
			}
			sNrTo, err := strconv.Atoi(value[to+3:])
			if err != nil {
				log.Fatal("Error to")
			}
			fmt.Println(sCount, "x", sNrFrom, "x", sNrTo)
			m := move{sCount, sNrFrom, sNrTo}
			moveSteps = append(moveSteps, m)
		} else {
			v := value
			pos := 0
			index := 0

			for {
				x := strings.IndexAny(v, "[")
				if x == -1 {
					break
				}
				index++
				pos += x + 1
				y := strings.IndexAny(v, "]")
				fmt.Println(v, x, y, pos)
				var found *stack
				for _, idx := range stackIndex {
					if idx.index == pos {
						found = idx
						break
					}
				}
				if found == nil {
					s := NewStack()
					si := &stack{pos, s}
					stackIndex = append(stackIndex, si)
					found = si
				}
				fmt.Println(v[x+1:y], x, pos)
				c := &crates{pos, v[x+1 : y]}
				found.st.Low(c)
				v = v[y+1:]
				pos += y - x
			}
		}
	}

	for _, m := range moveSteps {
		fi := stackIndex[m.from-1]
		ti := stackIndex[m.to-1]
		fmt.Println(m.count, fi.index, "->", ti.index)
		for i := 0; i < m.count; i++ {
			x, err := fi.st.Pop()
			if err != nil {
				log.Fatal("Error stack")
			}
			fmt.Println("Take", x)
			fmt.Println("lay", x)
			ti.st.Push(x)
		}
	}
	finalWord := ""
	for _, s := range stackIndex {
		x, err := s.st.Pop()
		if err != nil {
			log.Fatal("Error stac2")
		}
		c := x.(*crates)
		fmt.Println(s.index, c.v)
		finalWord += c.v
	}
	fmt.Println("Word:", finalWord)
}

func TestAdventGame5b(t *testing.T) {
	fInput, ferr := os.Open("files/input.5.save")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	moves := false
	stackIndex := make([]*stack, 0)
	moveSteps := make([]move, 0)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)

		if value == "" {
			moves = true
			sort.Slice(stackIndex, func(i, j int) bool {
				return stackIndex[i].index < stackIndex[j].index
			})
			for _, s := range stackIndex {
				fmt.Println(s.index, s.st.Size)
			}
			continue
		}
		if moves {
			from := strings.Index(value, "from")
			to := strings.Index(value, "to")
			sCount, err := strconv.Atoi(value[5 : from-1])
			if err != nil {
				log.Fatal("Error count >" + value[5:from] + "<" + err.Error())
			}
			sNrFrom, err := strconv.Atoi(value[from+5 : to-1])
			if err != nil {
				log.Fatal("Error from >" + value[from+5:to] + "<" + err.Error())
			}
			sNrTo, err := strconv.Atoi(value[to+3:])
			if err != nil {
				log.Fatal("Error to")
			}
			fmt.Println(sCount, "x", sNrFrom, "x", sNrTo)
			m := move{sCount, sNrFrom, sNrTo}
			moveSteps = append(moveSteps, m)
		} else {
			v := value
			pos := 0
			index := 0

			for {
				x := strings.IndexAny(v, "[")
				if x == -1 {
					break
				}
				index++
				pos += x + 1
				y := strings.IndexAny(v, "]")
				fmt.Println(v, x, y, pos)
				var found *stack
				for _, idx := range stackIndex {
					if idx.index == pos {
						found = idx
						break
					}
				}
				if found == nil {
					s := NewStack()
					si := &stack{pos, s}
					stackIndex = append(stackIndex, si)
					found = si
				}
				fmt.Println(v[x+1:y], x, pos)
				c := &crates{pos, v[x+1 : y]}
				found.st.Low(c)
				v = v[y+1:]
				pos += y - x
			}
		}
	}

	for _, m := range moveSteps {
		fi := stackIndex[m.from-1]
		ti := stackIndex[m.to-1]
		fmt.Println(m.count, fi.index, "->", ti.index)
		crane := make([]*crates, 0)
		for i := 0; i < m.count; i++ {
			x, err := fi.st.Pop()
			if err != nil {
				log.Fatal("Error stack")
			}
			fmt.Println("Take", x)
			crane = append(crane, x.(*crates))
		}
		for i := m.count; i > 0; i-- {
			c := crane[i-1]
			fmt.Println("lay", c)
			ti.st.Push(c)
		}
	}
	finalWord := ""
	for _, s := range stackIndex {
		x, err := s.st.Pop()
		if err != nil {
			log.Fatal("Error stac2")
		}
		c := x.(*crates)
		fmt.Println(s.index, c.v)
		finalWord += c.v
	}
	fmt.Println("Word:", finalWord)
}
