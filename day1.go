package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	testLabel := os.Getenv("Test")
	var test bool = true
	if testLabel == "" {
		test = false
	}
	data, err := filereader.Reader(test, 1)
	if err != nil {
		return
	}

	var eachData = strings.Split(data, "\n")
	eachData = append(eachData, "")
	finalData := []int{}
	sums := []int{}
	for _, d := range eachData {
		if d != "" {
			i, err := strconv.Atoi(d)
			if err != nil {
				fmt.Println("error while converting", err)
			}
			finalData = append(finalData, i)
		} else {
			singleSum := 0
			for _, cal := range finalData {
				singleSum += cal
			}
			sums = append(sums, singleSum)
			finalData = []int{}
		}
	}

	maxCal := 0
	for _, s := range sums {
		if s > maxCal {
			maxCal = s
		}
	}
	sort.Ints(sums)
	fmt.Println("here is highest", maxCal)
	fmt.Println("here is totalcal", sums[len(sums)-1]+sums[len(sums)-2]+sums[len(sums)-3])
}
