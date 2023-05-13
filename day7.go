package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	originalData []string
	currentLine  int
}

func main() {
	testLabel := os.Getenv("Test")
	var test bool = true
	if testLabel == "" {
		test = false
	}
	data, _ := filereader.Reader(test, 7)
	var eachData = strings.Split(data, "\n")
	eachData = append(eachData, "")
	tree := Tree{eachData, 0}
	tree.traverse()
}

func (tree *Tree) traverse() {
	dirStruct := map[string][]string{}
	currentDir := ""
	lastDir := []string{}
	for num, d := range tree.originalData {
		if d == "" {
			continue
		}
		tree.currentLine = num
		var cmds = strings.Split(d, " ")
		if strings.HasPrefix(d, "$ cd") {
			if strings.HasPrefix(d, "$ cd ..") {
				currentDir = lastDir[len(lastDir)-1]
				lastDir = lastDir[:len(lastDir)-1]
			} else {
				currentDir = cmds[2]
				lastDir = append(lastDir, currentDir)
			}
		} else {
			structure := strings.Split(d, " ")
			if strings.HasPrefix(d, "$ ls") {
				continue
			}
			running := strings.Join(lastDir, ",")
			if strings.Contains(structure[0], "dir") {
				dirStruct[running] = append(dirStruct[running], strings.Join(lastDir, ",")+","+structure[1])
			} else {
				dirStruct[running] = append(dirStruct[running], structure[0])
			}
		}
	}
	pureElems := map[string]bool{}
	onlyFileSizesinRoot := 0
	for _, f := range dirStruct["/"] {
		num, err := strconv.Atoi(f)
		if err == nil {
			onlyFileSizesinRoot += num
		}
	}
	for {
		if len(pureElems) == len(dirStruct) {
			break
		}
		for key, files := range dirStruct {
			pure := true
			for _, f := range files {
				_, err := strconv.Atoi(f)
				if err != nil {
					pure = false
				}
			}
			if pure {
				sum := 0
				for _, f := range files {
					val, err := strconv.Atoi(f)
					if err == nil {
						sum += val
					}
				}
				dirStruct[key] = []string{strconv.Itoa(sum)}
				for kk, val := range dirStruct {
					for _, purity := range val {
						if purity == key {
							removed := remove(val, purity)
							dirStruct[kk] = removed
							dirStruct[kk] = append(dirStruct[kk], strconv.Itoa(sum))
						}
					}
				}
				pureElems[key] = true
			}
		}
	}
	firstSum := 0
	slashSize := 0
	for _, files := range dirStruct {
		num, err := strconv.Atoi(files[0])
		if err == nil && num < 100000 {
			firstSum += num
		}
	}
	fmt.Println("Sum of directories with files/dirs exceeding 100000 is:", firstSum)
	slashSize += onlyFileSizesinRoot

	/*
		Second Part
	*/

	tree.FindDirectoryToDelete(dirStruct)
}

func (tree *Tree) FindDirectoryToDelete(dir map[string][]string) {
	totalConsumed, _ := strconv.Atoi(dir["/"][0])
	unused := 70000000 - totalConsumed
	selectedForDeletion := []int{}
	for d, files := range dir {
		if d == "/" {
			continue
		}
		num, err := strconv.Atoi(files[0])
		if err == nil && unused+num > 30000000 {
			selectedForDeletion = append(selectedForDeletion, num)
		}
	}
	smallest := 1010101010101
	for _, small := range selectedForDeletion {
		if small < smallest {
			smallest = small
		}
	}
	fmt.Println("Size of the smallest directory to delete", smallest)
}

func remove(s []string, what string) []string {
	for i, v := range s {
		if v == what {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}
