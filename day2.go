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
	data, err := filereader.Reader(test, 2)
	if err != nil {
		return
	}
	var winmap = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
		"C": "X",
		"B": "Z",
		"A": "Y",
	}
	var drawmap = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	var lostmap = map[string]string{
		"X": "B",
		"Y": "C",
		"Z": "A",
		"B": "X",
		"C": "Y",
		"A": "Z",
	}
	var pointMap = map[string]int{
		"X": 1, //Rock
		"Y": 2, //Paper
		"Z": 3, //Scissors
		"A": 1, //Rock
		"B": 2, //Paper
		"C": 3, //Scissors
	}
	var eachData = strings.Split(data, "\n")
	var score int
	var score2 int
	for _, d := range eachData {
		var roundscore int
		if d == "" {
			continue
		}
		// fmt.Println("each", d)
		var round = strings.Split(d, " ")
		// fmt.Println(round[0])
		if winmap[round[1]] == round[0] {
			fmt.Println("Won Round", round[0], round[1])
			score += (pointMap[round[1]] + 6)
		} else if drawmap[round[1]] == round[0] {
			fmt.Println("Draw Round", round[0], round[1])
			score += (pointMap[round[1]] + 3)
		} else {
			fmt.Println("Lost Round", round[0], round[1])
			score += (pointMap[round[1]] + 0)
		}
		if round[1] == "X" {
			roundscore = pointMap[lostmap[round[0]]]
			fmt.Println("Lost", roundscore)
			score2 += pointMap[lostmap[round[0]]]
		} else if round[1] == "Y" {
			roundscore = pointMap[drawmap[round[0]]] + 3
			fmt.Println("Tie", roundscore)
			score2 += roundscore
		} else {
			roundscore = pointMap[winmap[round[0]]] + 6
			fmt.Println("Won", roundscore)
			score2 += (pointMap[winmap[round[0]]] + 6)
		}
	}
	fmt.Println("final score is", score)
	fmt.Println("final score 2 is", score2)
}
