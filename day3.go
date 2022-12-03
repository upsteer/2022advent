package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"strings"
)

var (
	alphabet    string = "abcdefghijklmnopqrstuvwxyz"
	bigAlphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	testLabel := os.Getenv("Test")
	var test bool = true
	if testLabel == "" {
		test = false
	}
	data, err := filereader.Reader(test, 3)
	if err != nil {
		return
	}
	var eachData = strings.Split(data, "\n")
	score, _ := calculate(eachData)
	fmt.Println("first", score)
	var score2 int
	for i := 0; i < len(eachData)-1; i += 3 {
		found1 := findCommononThree([]string{eachData[i], eachData[i+1], eachData[i+2]})
		if strings.Index(alphabet, found1) > -1 {
			score2 += strings.Index(alphabet, found1) + 1
		} else {
			score2 += strings.Index(bigAlphabet, found1) + 1
		}
	}
	fmt.Println("second", score2)
}

func findCommononThree(threegroup []string) string {
	var found string
	for i := 0; i < len(threegroup[0]); i++ {
		current := string(threegroup[0][i])
		for k := 0; k < len(threegroup[1]); k++ {
			current2 := string(threegroup[1][k])
			if current == current2 {
				for j := 0; j < len(threegroup[2]); j++ {
					current3 := string(threegroup[2][j])
					if current2 == current3 {
						found = current2
						break
					}
				}
			}
		}
	}
	return found
}

func calculate(eachData []string) (int, string) {
	var score int
	var common string
	for _, d := range eachData {
		if d == "" {
			continue
		}
		lenstr := len(string(d))
		secondcompartment := string(d[lenstr/2:])
		for i := 0; i < lenstr/2; i++ {
			if strings.Contains(secondcompartment, string(d[i])) {
				common = string(d[i])
				if strings.Index(alphabet, string(d[i])) > -1 {
					score += (strings.Index(alphabet, string(d[i])) + 1)
				} else {
					score += (strings.Index(bigAlphabet, string(d[i])) + 1)
				}
				break
			}
		}
	}
	return score, common
}
