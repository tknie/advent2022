package advent

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

type fstype byte

const (
	filesystem fstype = iota
	directory
)

type fs struct {
	name   string
	ty     fstype
	size   int
	sub    []*fs
	parent *fs
}

var fsList = make([]*fs, 0)

func TestAdventGame7a(t *testing.T) {
	fInput, ferr := os.Open("files/input.7.test")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	var cfs *fs
	var parent *fs
	list := false
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		switch {
		case strings.HasPrefix(value, "$ cd"):
			loc := value[5:]
			fmt.Printf("cd to <%s>\n", loc)
			switch loc {
			case "/":
				if parent != nil {
					cfs = parent
				} else {
					parent = &fs{ty: directory, name: loc}
					fsList = append(fsList, parent)
					cfs = parent
				}
			case "..":
				cfs = cfs.parent
			default:
				f := false
				for _, sb := range cfs.sub {
					if sb.name == loc {
						cfs = sb
						f = true
						break
					}
				}
				if !f {
					nfs := &fs{ty: directory, name: loc, parent: cfs,
						sub: make([]*fs, 0)}
					fsList = append(fsList, nfs)
					cfs = nfs
					fmt.Println("New current", loc)
				}
			}
			list = false
		case value == "$ ls":
			list = true
		default:
			if !list {
				log.Fatal("Errror unknown parser:" + value)
			}
			ds := strings.Split(value, " ")
			if ds[0] == "dir" {
				dfs := &fs{ty: directory, name: ds[1],
					parent: cfs, sub: make([]*fs, 0)}
				cfs.sub = append(cfs.sub, dfs)
				fsList = append(fsList, dfs)
			} else {
				ffs := &fs{ty: filesystem, name: ds[1]}
				cfs.sub = append(cfs.sub, ffs)
				sz, err := strconv.Atoi(ds[0])
				if err != nil {
					log.Fatal("Error working on size:" + ds[0])
				}
				ffs.size = sz
				cfs.size += sz
			}
		}
	}
	refreshD(parent)
	fmt.Println("Dump:", parent.name)
	dump(parent, 0)
	fmt.Println("Max:", calculateMax(parent, 100000))
}

func refreshD(f *fs) {
	for _, sf := range f.sub {
		if sf.ty == directory {
			refreshD(sf)
		}
	}

	f.size = 0
	for _, sf := range f.sub {
		f.size += sf.size
	}

}

func dump(f *fs, deep int) {

	ps := strconv.Itoa(deep) + strings.Repeat(" ", deep)
	if f.ty == directory {
		ps += "d"
	} else {
		ps += "f"
	}
	fmt.Println(ps, f.name, f.size)
	for _, sf := range f.sub {
		dump(sf, deep+1)
	}
}

func calculateMax(f *fs, max int) int {
	total := 0
	for _, sf := range f.sub {
		if sf.ty == directory {
			total += calculateMax(sf, max)
		}
	}
	if f.ty == directory && f.size < max {
		total += f.size
	}
	return total
}

func calculateSmallest(f *fs, max int) int {
	if f.size < max {
		return math.MaxInt
	}
	small := f.size
	for _, sf := range f.sub {
		if sf.ty == directory {
			subsmall := calculateSmallest(sf, max)
			if subsmall >= max && subsmall < small {
				small = subsmall
			}
			if sf.size >= max && sf.size < small {
				small = sf.size
			}
		}
	}
	return small
}

func TestAdventGame7b(t *testing.T) {
	fInput, ferr := os.Open("files/input.7.save")
	if ferr != nil {
		fmt.Println("Error reading input", ferr)
		return
	}
	defer fInput.Close()
	fileScanner := bufio.NewScanner(fInput)
	fileScanner.Split(bufio.ScanLines)
	var cfs *fs
	var parent *fs
	list := false
	for fileScanner.Scan() {
		value := fileScanner.Text()
		fmt.Println(value)
		switch {
		case strings.HasPrefix(value, "$ cd"):
			loc := value[5:]
			fmt.Printf("cd to <%s>\n", loc)
			switch loc {
			case "/":
				if parent != nil {
					cfs = parent
				} else {
					parent = &fs{ty: directory, name: loc}
					fsList = append(fsList, parent)
					cfs = parent
				}
			case "..":
				cfs = cfs.parent
			default:
				f := false
				for _, sb := range cfs.sub {
					if sb.name == loc {
						cfs = sb
						f = true
						break
					}
				}
				if !f {
					nfs := &fs{ty: directory, name: loc, parent: cfs,
						sub: make([]*fs, 0)}
					fsList = append(fsList, nfs)
					cfs = nfs
					fmt.Println("New current", loc)
				}
			}
			list = false
		case value == "$ ls":
			list = true
		default:
			if !list {
				log.Fatal("Errror unknown parser:" + value)
			}
			ds := strings.Split(value, " ")
			if ds[0] == "dir" {
				dfs := &fs{ty: directory, name: ds[1],
					parent: cfs, sub: make([]*fs, 0)}
				cfs.sub = append(cfs.sub, dfs)
				fsList = append(fsList, dfs)
			} else {
				ffs := &fs{ty: filesystem, name: ds[1]}
				cfs.sub = append(cfs.sub, ffs)
				fsList = append(fsList, ffs)
				sz, err := strconv.Atoi(ds[0])
				if err != nil {
					log.Fatal("Error working on size:" + ds[0])
				}
				ffs.size = sz
				cfs.size += sz
			}
		}
	}
	refreshD(parent)
	fmt.Println("Dump:", parent.name)
	dump(parent, 0)
	freeSize := 70000000 - parent.size
	neededSize := 30000000 - freeSize
	fmt.Println(freeSize)
	fmt.Println("Find smallest in tree:", calculateSmallest(parent, neededSize))
	small := 0
	for _, fs := range fsList {
		if fs.size >= neededSize {
			if small == 0 || small > fs.size {
				small = fs.size
			}
		}
	}
	fmt.Println("Find smallest in list:", small)
	fmt.Println("Free Size            :", freeSize)
	fmt.Println("At least delete      :", neededSize)

}
