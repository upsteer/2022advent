package main

import (
	"advent-2022/filereader"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	xstartingPos     = 0
	ystartingPos     = 99
	xheadstartingPos = 0
	yheadstartingPos = 99
	xtailstartingPos = 0
	ytailstartingPos = 99
	tailMovement     = ""
)

func main() {
	testLabel := os.Getenv("Test")
	var test bool = true

	if testLabel == "" {
		test = false
	}
	data, err := filereader.Reader(test, 9)
	if err != nil {
		return
	}

	var eachData = strings.Split(data, "\n")
	eachData = append(eachData, "")
	visited := map[string]bool{}
	for _, d := range eachData {
		if d != "" {
			cmds := strings.Split(d, " ")
			i, _ := strconv.Atoi(cmds[1])
			fmt.Println("move", cmds[0], i)
			for j := 0; j < i; j++ {
				moveHeadToDirection(cmds[0])
				xstr := strconv.Itoa(xtailstartingPos)
				ystr := strconv.Itoa(ytailstartingPos)
				// xxstr := strconv.Itoa(xheadstartingPos)
				// yystr := strconv.Itoa(yheadstartingPos)
				// fmt.Println("head", xxstr, yystr)
				fmt.Println("tail", xstr, ystr)
				visited[strings.Join([]string{xstr, ystr}, ",")] = true
			}
			tailMovement = ""
		}
	}
	fmt.Println(visited)
	fmt.Println("visited", len(visited))
}
func diff(a int, b int) int {
	return a - b
}
func moveHeadToDirection(direction string) {
	switch direction {
	case "L":
		xheadstartingPos -= 1
	case "R":
		xheadstartingPos += 1
	case "U":
		yheadstartingPos -= 1
	case "D":
		yheadstartingPos += 1
	}
	if len(tailMovement) > 0 {
		if shouldMove() {
			if len(tailMovement) > 1 {
				moveDiagonal(tailMovement)
			} else {
				moveTailToDirection(tailMovement)
			}
		}
	}
	tailMovement = whereToMove()
	// fmt.Println("tailMovement", tailMovement)
}

func shouldMove() bool {
	xdiff := Abs(diff(xheadstartingPos, xtailstartingPos))
	ydiff := Abs(diff(yheadstartingPos, ytailstartingPos))
	if xdiff > 1 || ydiff > 1 {
		return true
	}
	if xdiff == 0 || ydiff == 0 && (xdiff > 0 || ydiff > 0) {
		return false
	}
	return false
}

func moveTailToDirection(direction string) {
	switch direction {
	case "L":
		xtailstartingPos -= 1
	case "R":
		xtailstartingPos += 1
	case "U":
		ytailstartingPos -= 1
	case "D":
		ytailstartingPos += 1
	}
}
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func whereToMove() string {
	xdiff := diff(xheadstartingPos, xtailstartingPos)
	ydiff := diff(yheadstartingPos, ytailstartingPos)
	movement := ""
	if Abs(xdiff) > 0 && Abs(ydiff) > 0 {
		movement = determineTailMovement(xheadstartingPos, yheadstartingPos, xtailstartingPos, ytailstartingPos)
		return movement
	}
	if xdiff > 0 {
		return "R"
	} else if xdiff < 0 {
		return "L"
	}
	if ydiff > 0 {
		return "D"
	} else if ydiff < 0 {
		return "U"
	}
	return movement
}
func determineTailMovement(xhead int, yhead int, xtail int, ytail int) string {
	dir := ""
	if yhead-ytail > 0 {
		dir += "D"
	} else {
		dir += "U"
	}
	if xhead-xtail > 0 {
		dir += "R"
	} else {
		dir += "L"
	}
	return dir
}
func moveDiagonal(direction string) {
	switch direction {
	case "UL":
		xtailstartingPos -= 1
		ytailstartingPos -= 1
	case "UR":
		xtailstartingPos += 1
		ytailstartingPos -= 1
	case "DL":
		xtailstartingPos -= 1
		ytailstartingPos += 1
	case "DR":
		xtailstartingPos += 1
		ytailstartingPos += 1
	}
}
