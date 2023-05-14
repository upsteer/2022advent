package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

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

				}
				rowData = append(rowData, i)
			}
			finalData = append(finalData, rowData)
		}
	}
	visibleTreesFromEdge := 0
	scenicList := []int{}
	maxSceneScore := MinInt
	for y, elemy := range finalData {
		if !(y == 0 || y == len(finalData)-1) {
			for x, _ := range elemy {
				if !(x == 0 || x == len(elemy)-1) {
					visible, scenicQuotient := checkVisibility(finalData, y, x, len(finalData)-1, len(elemy)-1)
					if scenicQuotient > maxSceneScore {
						maxSceneScore = scenicQuotient
					}
					scenicList = append(scenicList, scenicQuotient)
					if visible {
						visibleTreesFromEdge += 1
						continue
					}
				}
			}
		}
	}
	visibleTreesFromEdge += (len(finalData)*2 + (len(finalData)-2)*2)
	fmt.Println("Trees that are visible from the edge", visibleTreesFromEdge)
	fmt.Println("Best Scenic Quotient", maxSceneScore)
}

func checkVisibility(forest [][]int, ypos int, xpos int, ymax int, xmax int) (bool, int) {
	northVisibility := true
	southVisibility := true
	eastVisibility := true
	westVisibility := true
	scenicQuotient := 1
	tree := forest[ypos][xpos]
	for n := ypos - 1; n >= 0; n-- {
		north := forest[n][xpos]
		if north >= tree {
			scenicQuotient *= (ypos - n)
			northVisibility = false
			break
		}
	}
	if northVisibility {
		scenicQuotient *= ypos
	}
	for s := ypos + 1; s <= ymax; s++ {
		south := forest[s][xpos]
		if south >= tree {
			scenicQuotient *= (s - ypos)
			southVisibility = false
			break
		}
	}
	if southVisibility {
		scenicQuotient *= (ymax - ypos)
	}
	for e := xpos + 1; e <= xmax; e++ {
		east := forest[ypos][e]
		if east >= tree {
			scenicQuotient *= (e - xpos)
			eastVisibility = false
			break
		}
	}
	if eastVisibility {
		scenicQuotient *= (xmax - xpos)
	}
	for w := xpos - 1; w >= 0; w-- {
		west := forest[ypos][w]
		if west >= tree {
			scenicQuotient *= (xpos - w)
			westVisibility = false
			break
		}
	}
	if westVisibility {
		scenicQuotient *= xpos
	}
	if northVisibility || southVisibility || eastVisibility || westVisibility {
		return true, scenicQuotient
	}
	return false, scenicQuotient
}
