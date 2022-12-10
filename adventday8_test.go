package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestAdventGame8a(t *testing.T) {
	fInput, ferr := os.Open("files/input.8.save")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	vba := make([][][]byte, 0)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)

		v := []byte(value)
		vb := make([][]byte, 0)
		for _, v := range v {
			x, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatal("Error atoi")
			}
			vb = append(vb, []byte{byte(x), 0})
			fmt.Println(vb)
		}
		vba = append(vba, vb)
	}
	for j := 0; j < len(vba[0]); j++ {
		for i := 0; i < len(vba); i++ {
			// if j == 0 || i == 0 || j == len(vba[0])-1 || i == len(vba)-1 {
			// 	vba[j][i][1] = 1
			// } else {
			fmt.Println("Work on ", vba[j][i], j, i)
			m := byte(0)
			for k := 0; k < j; k++ {
				switch {
				case m == 0:
					m = vba[k][i][0]
					if vba[k][i][1] == 0 {
						vba[k][i][1] = 1
						fmt.Println("A", m, k, i, vba[k][i])
					} else {
						fmt.Println("B", m, k, i, vba[k][i])
					}
				case vba[k][i][1] == 0:
					if m > 0 && vba[k][i][0] > m {
						if vba[k][i][1] == 0 {
							vba[k][i][1] = 1
							fmt.Println("C", m, k, i, vba[k][i])
						} else {
							fmt.Println("D", m, k, i, vba[k][i])
						}
						m = vba[k][i][0]
					}
				}
			}
			m = byte(0)
			for k := len(vba[0]) - 1; k > j; k-- {
				switch {
				case m == 0:
					m = vba[k][i][0]
					if vba[k][i][1] == 0 {
						vba[k][i][1] = 1
						fmt.Println("E", m, k, i, vba[k][i])
					} else {
						fmt.Println("F", m, k, i, vba[k][i])
					}
				case m > 0 && vba[k][i][0] > m:
					if vba[k][i][1] == 0 {
						vba[k][i][1] = 1
						fmt.Println("G", m, k, i, vba[k][i])
					} else {
						fmt.Println("H", m, k, i, vba[k][i])
					}
					m = vba[k][i][0]
				default:
					fmt.Println("HI", m, k, i, vba[k][i])
				}
			}
			m = byte(0)

			for k := 0; k < i; k++ {
				switch {
				case m == 0:
					m = vba[j][k][0]
					if vba[j][k][1] == 0 {
						vba[j][k][1] = 1
						fmt.Println("I", m, j, k, vba[j][k])
					} else {
						fmt.Println("J", m, j, k, vba[j][k])
					}
				case m > 0 && vba[j][k][0] > m:
					if vba[j][k][1] == 0 {
						vba[j][k][1] = 1
						fmt.Println("K", m, j, k, vba[j][k])
					} else {
						fmt.Println("L", m, j, k, vba[j][k])
					}
					m = vba[j][k][0]
				default:
					fmt.Println("OO", m, j, k, vba[j][k])
				}
			}
			m = byte(0)
			for k := len(vba) - 1; k > i; k-- {
				switch {
				case m == 0:
					m = vba[j][k][0]
					if vba[j][k][1] == 0 {
						vba[j][k][1] = 1
						fmt.Println("M", m, j, k, vba[j][k])
					} else {
						fmt.Println("N", m, j, k, vba[j][k])
					}
				case m > 0 && vba[j][k][0] > m:
					if vba[j][k][1] == 0 {
						vba[j][k][1] = 1
						fmt.Println("O", m, j, k, vba[j][k])
					} else {
						fmt.Println("P", m, j, k, vba[j][k])
					}
					m = vba[j][k][0]

				default:
					fmt.Println("P", m, j, k, vba[j][k])
				}
			}
		}
		//}
	}
	check(vba)

	visible := 0
	fmt.Println("Dump")
	for x, vb := range vba {
		fmt.Println(x, vb)
		for _, v := range vb {
			if v[1] == 1 {
				visible++
			}
		}
	}
	fmt.Println(visible)
	fmt.Println(len(vba), len(vba[0]))
}

func check(vba [][][]byte) {
	vs := 0
	maxscore := 0
	for j := 0; j < len(vba); j++ {
		for i := 0; i < len(vba[j]); i++ {
			var b [4]bool
			b[0] = true
			b[1] = true
			b[2] = true
			b[3] = true
			score := 0
			m1 := 0
			for x := i - 1; x >= 0; x-- {
				m1++
				if vba[j][x][0] >= vba[j][i][0] {
					b[0] = false
					break
				}
			}
			m2 := 0
			for x := i + 1; x < len(vba[j]); x++ {
				m2++
				if vba[j][x][0] >= vba[j][i][0] {
					b[1] = false
					break
				}
			}
			m3 := 0
			for x := j - 1; x >= 0; x-- {
				m3++
				if vba[x][i][0] >= vba[j][i][0] {
					b[2] = false
					break
				}
			}
			m4 := 0
			for x := j + 1; x < len(vba); x++ {
				m4++
				if vba[x][i][0] >= vba[j][i][0] {
					b[3] = false
					break
				}
			}
			score = m1 * m2 * m3 * m4
			if score > maxscore {
				maxscore = score
			}
			if !b[0] && !b[1] && !b[2] && !b[3] {
				if i == 0 || j == 0 || i == len(vba[j])-1 || j == len(vba)-1 {
					vba[j][i][1] = 1
					vs++
				} else {
					vba[j][i][1] = 0
				}
			} else {
				vba[j][i][1] = 1
				vs++
			}
			fmt.Println("Score", i, j, score, vba[j][i])
		}
	}
	fmt.Println("Max", maxscore)
	fmt.Println("V2", vs)
}

func TestAdventGame8b(t *testing.T) {
	fInput, ferr := os.Open("files/input.8.test")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	vba := make([][]byte, 0)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		vb := []byte(value)
		vba = append(vba, vb)
	}
	for i := 0; i < len(vba); i++ {
		b := vba[i]
		fmt.Println(b)
	}
	finalWord := ""
	fmt.Println("Word:", finalWord)
}
