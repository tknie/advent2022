package advent

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"testing"
)

type sortRunes []byte

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) []byte {
	r := []byte(s)
	sort.Sort(sortRunes(r))
	return r
}

func TestAdventGame3(t *testing.T) {
	fInput, ferr := os.Open("files/input.3")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	total := 0
	c := byte('a')
	d := byte('A')
	e := byte('z')
	f := byte('Z')
	fmt.Println(c, e, d, f)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		a := value[:len(value)/2]
		b := value[len(value)/2:]
		aCA := SortString(a)
		bCA := SortString(b)
		found := byte(0)
		for _, aC := range aCA {
			for _, bC := range bCA {
				if aC == bC {
					found = aC
					break
				}
			}
			if found != 0 {
				break
			}
		}
		fmt.Printf("Found %d %c\n", found, rune(found))
		if found >= c && found <= e {
			prio := found - c + 1
			fmt.Printf("Prio %d for %c\n", prio, found)
			total += int(prio)
		}
		if found >= d && found <= f {
			prio := found - d + 27
			fmt.Printf("Prio %d for %c\n", prio, found)
			total += int(prio)
		}
	}
	fmt.Println("Prio all:", total)
}

func TestAdventGame3b(t *testing.T) {
	fInput, ferr := os.Open("files/input.3")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	total := 0
	c := byte('a')
	d := byte('A')
	e := byte('z')
	f := byte('Z')
	fmt.Println(c, e, d, f)
	in := make([]string, 3)
	cindex := 0
	for fileScanner.Scan() {
		value := fileScanner.Text()
		i := cindex % 3
		in[i] = value
		cindex++
		fmt.Println(i, value)
		if i == 2 {
			aCA := SortString(in[0])
			bCA := SortString(in[1])
			cCA := SortString(in[2])
			found := byte(0)
			for _, aC := range aCA {
				for _, bC := range bCA {
					if aC == bC {
						for _, cC := range cCA {
							if aC == cC {
								found = aC
								break
							}
						}
					}
					if found != 0 {
						break
					}
				}
				if found != 0 {
					break
				}
			}
			fmt.Printf("Found %d %c\n", found, rune(found))
			if found >= c && found <= e {
				prio := found - c + 1
				fmt.Printf("Prio %d for %c\n", prio, found)
				total += int(prio)
			}
			if found >= d && found <= f {
				prio := found - d + 27
				fmt.Printf("Prio %d for %c\n", prio, found)
				total += int(prio)
			}
		}
	}
	fmt.Println("Prio all:", total)
}
