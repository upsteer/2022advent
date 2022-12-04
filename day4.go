package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	testLabel := os.Getenv("Test")
	var test bool = true
	if testLabel == "" {
		test = false
	}
	data, err := filereader.Reader(test, 4)
	if err != nil {
		return
	}
	var eachData = strings.Split(data, "\n")
	counter := 0
	counter2 := 0
	for _, d := range eachData {
		if d == "" {
			continue
		}
		var sections = strings.Split(d, ",")
		lengths := []int{}
		arrays := [][]int32{}
		for _, sec := range sections {
			var ends = strings.Split(sec, "-")
			lengths = append(lengths, (toint(ends[1]) - toint(ends[0]) + 1))
			arrays = append(arrays, makenaturalarray(toint(ends[0]), toint(ends[1])))
		}
		shortest := 0
		if len(arrays[1]) < len(arrays[0]) {
			shortest = 1
		}
		if subset(arrays[shortest], arrays[1-shortest]) {
			counter += 1
		}
		for _, fff := range arrays[0] {
			breakCondition := false
			for _, sss := range arrays[1] {
				if sss == fff {
					counter2 += 1
					breakCondition = true
					break
				}
			}
			if breakCondition {
				break
			}
		}
	}
	fmt.Println("first", counter)
	fmt.Println("second", counter2)
}
func toint(numstr string) int {
	i, err := strconv.Atoi(numstr)
	if err != nil {
		fmt.Println("error while converting", err)
	}
	return i
}

func makenaturalarray(from int, till int) []int32 {
	make := []int32{}
	for i := from; i <= till; i++ {
		make = append(make, int32(i))
	}
	return make
}

func subset(first, second []int32) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[int(value)] += 1
	}

	for _, value := range first {
		if count, found := set[int(value)]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[int(value)] = count - 1
		}
	}

	return true
}
