package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"strings"
)

func main() {
	testLabel := os.Getenv("Test")
	var test bool = true
	if testLabel == "" {
		test = false
	}
	data, err := filereader.Reader(test, 6)
	if err != nil {
		return
	}
	var eachData = strings.Split(data, "")
	for i := 0; i < len(eachData); i++ {
		if unique(eachData[i : i+4]) {
			fmt.Println("answer", i+4)
			break
		}
	}
	for i := 0; i < len(eachData); i++ {
		if unique(eachData[i : i+14]) {
			fmt.Println("answer2", i+14)
			break
		}
	}
}

func unique(listofchars []string) bool {
	isunique := true
	for i, item := range listofchars {
		for j := i + 1; j < len(listofchars); j++ {
			if listofchars[j] == item {
				isunique = false
				break
			}
		}
	}
	return isunique
}
