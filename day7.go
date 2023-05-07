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
			fmt.Println("curr", lastDir)
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
				_, err := strconv.Atoi(f)
				if err != nil {
					pure = false
					file := remove(files, index)
					file = append(file, dirStruct[f]...)
					dirStruct[key] = file
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

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
