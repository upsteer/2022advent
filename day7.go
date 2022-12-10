package main

import (
	"advent-2022/filereader"
	"encoding/json"
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
	currentDir := "/"
	lastDir := []string{}
	for num, d := range tree.originalData {
		if d == "" {
			continue
		}
		tree.currentLine = num
		var cmds = strings.Split(d, " ")
		if strings.HasPrefix(d, "$ ls") {
			// calculateFileSizes()
		} else if strings.HasPrefix(d, "$ cd") {
			if strings.HasPrefix(d, "$ cd ..") {
				// fmt.Println(num, d)
				currentDir = lastDir[len(lastDir)-1]
				lastDir = lastDir[:len(lastDir)-1]
			} else {
				lastDir = append(lastDir, currentDir)
				currentDir = cmds[2]
				// fmt.Println(num, d)
			}
		} else {
			structure := strings.Split(d, " ")
			if strings.Contains(structure[0], "dir") {
				dirStruct[currentDir] = append(dirStruct[currentDir], d)
			} else {
				dirStruct[currentDir] = append(dirStruct[currentDir], structure[0])
			}
		}
	}
	jsStr, _ := json.Marshal(dirStruct)
	fmt.Println(string(jsStr))
	pureElems := map[string]bool{}
	for {
		if len(pureElems) == len(dirStruct) {
			break
		}
		for key, files := range dirStruct {
			pure := true
			for index, f := range files {
				childs := strings.Split(f, " ")
				if len(childs) > 1 {
					pure = false
					if len(dirStruct[childs[1]]) == 1 {
						dirStruct[key][index] = dirStruct[childs[1]][0]
					}
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
				pureElems[key] = true
			}
		}

	}
	firstSum := 0
	for _, files := range dirStruct {
		num, err := strconv.Atoi(files[0])
		if err == nil && num < 100000 {
			firstSum += num
		}
	}
	fmt.Println("first", firstSum)
}
