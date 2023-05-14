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
	data, err := filereader.Reader(test, 8)
	if err != nil {
		return
	}

	eachData := strings.Split(data, "\n")
	eachData = append(eachData, "")
	finalData := [][]int{}
	for _, d := range eachData {
		if d != "" {
			row := strings.Split(d, "")
			rowData := []int{}
			for _, tree := range row {
				i, err := strconv.Atoi(tree)
				if err != nil {
					fmt.Println("error while converting", err)
				}
				rowData = append(rowData, i)
			}
			finalData = append(finalData, rowData)
		}
	}
	visibleTreesFromEdge := 0
	for y, elemy := range finalData {
		if !(y == 0 || y == len(finalData)-1) {
			for x, _ := range elemy {
				if !(x == 0 || x == len(elemy)-1) {
					if checkVisibility(finalData, y, x, len(finalData)-1, len(elemy)-1) {
						visibleTreesFromEdge += 1
						continue
					}
				}
			}
		}
	}
	visibleTreesFromEdge += (len(finalData)*2 + (len(finalData)-2)*2)
	fmt.Println("Trees that are visible from the edge", visibleTreesFromEdge)
}

func checkVisibility(forest [][]int, ypos int, xpos int, ymax int, xmax int) bool {
	northVisibility := true
	southVisibility := true
	eastVisibility := true
	westVisibility := true
	tree := forest[ypos][xpos]
	for n := ypos - 1; n >= 0; n-- {
		north := forest[n][xpos]
		if north >= tree {
			northVisibility = false
		}
	}
	for s := ypos + 1; s <= ymax; s++ {
		south := forest[s][xpos]
		if south >= tree {
			southVisibility = false
		}
	}
	for e := xpos + 1; e <= xmax; e++ {
		east := forest[ypos][e]
		if east >= tree {
			eastVisibility = false
		}
	}
	for w := xpos - 1; w >= 0; w-- {
		west := forest[ypos][w]
		if west >= tree {
			westVisibility = false
		}
	}
	if northVisibility || southVisibility || eastVisibility || westVisibility {
		return true
	}
	return false
}
