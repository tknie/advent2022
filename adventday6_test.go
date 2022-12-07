package advent

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestAdventGame6a(t *testing.T) {
	fInput, ferr := os.Open("files/input.6.save")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		values := []byte(value)
		//fmt.Println("Sources:", values)
		seq := make([]byte, 4)
		l := len(values)
		found := false
		for i := range values {
			if i < l-3 {
				copy(seq[:], value[i:i+4])
				// fmt.Println("Check:", seq)
				found = true
				for i := range seq {
					for j := i + 1; j < len(seq); j++ {
						if seq[i] == seq[j] {
							found = false
							break
						}
					}
					if !found {
						break
					}
				}
			}
			if found {
				fmt.Println("Found", i+1+3, string(seq))
				break
			}
		}
	}
	finalWord := ""
	fmt.Println("Word:", finalWord)
}

func TestAdventGame6b(t *testing.T) {
	fInput, ferr := os.Open("files/input.6.save")
	if ferr != nil {
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		values := []byte(value)
		// fmt.Println("Sources:", values)
		seq := make([]byte, 14)
		l := len(values)
		found := false
		for i := range values {
			if i < l-3 {
				copy(seq[:], value[i:i+16])
				//fmt.Println("Check:", seq)
				found = true
				for i := range seq {
					for j := i + 1; j < len(seq); j++ {
						if seq[i] == seq[j] {
							found = false
							break
						}
					}
					if !found {
						break
					}
				}
			}
			if found {
				fmt.Println("Found", i+1+13, string(seq))
				break
			}
		}
	}
	finalWord := ""
	fmt.Println("Word:", finalWord)
}
