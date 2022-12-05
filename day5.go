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
	data, err := filereader.Reader(test, 5)
	if err != nil {
		return
	}
	var eachData = strings.Split(data, "\n")
	movement := []string{}
	stackMap := map[string][]string{}
	for linenum, d := range eachData {
		if d == "" {
			continue
		}
		if strings.HasPrefix(d, "move") {
			movement = append(movement, d)
		} else if strings.HasPrefix(d, " 1") {
			var eachStacks = strings.Split(d, "")
			for index, stack := range eachStacks {
				if stack == " " {
					continue
				}
				for i := linenum - 1; i >= 0; i-- {
					containers := string(eachData[i][index])
					if containers == " " {
						continue
					}
					stackMap[stack] = append(stackMap[stack], containers)
				}
			}
		}
	}
	for _, move := range movement {
		var each = strings.Split(move, " ")
		nums := []int{}
		for _, num := range each {
			i, err := strconv.Atoi(num)
			if err != nil {
				continue
			}
			nums = append(nums, i)
		}
		//For Part 1, just uncomment this if clause and comment the else clause
		if nums[0] == 1 {
			for j := 0; j < nums[0]; j++ {
				pos := strconv.Itoa(nums[1])
				popped := stackMap[pos][len(stackMap[pos])-1]
				stackMap[pos] = stackMap[pos][:len(stackMap[pos])-1]
				stackMap[strconv.Itoa(nums[2])] = append(stackMap[strconv.Itoa(nums[2])], popped)
			}
		} else {
			poppedArr := []string{}
			for j := 0; j < nums[0]; j++ {
				pos := strconv.Itoa(nums[1])
				poppedArr = append(poppedArr, stackMap[pos][len(stackMap[pos])-1])
				stackMap[pos] = stackMap[pos][:len(stackMap[pos])-1]
			}
			for i, j := 0, len(poppedArr)-1; i < j; i, j = i+1, j-1 {
				poppedArr[i], poppedArr[j] = poppedArr[j], poppedArr[i]
			}
			stackMap[strconv.Itoa(nums[2])] = append(stackMap[strconv.Itoa(nums[2])], poppedArr...)
		}
	}
	finale := ""
	for c := 1; c < len(stackMap)+1; c++ {
		finale += stackMap[strconv.Itoa(c)][len(stackMap[strconv.Itoa(c)])-1]
	}
	fmt.Println("answer", finale)
}
