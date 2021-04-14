package main

import (
	"fmt"
	"math"
)

type asteroid struct {
	name   byte
	time1w int
	time2w int
	time3w int
	time1h int
	time2h int
	time3h int
}

var asteroidList []asteroid

func checkIfExist(ast byte) int {
	for i := 0; i < len(asteroidList); i++ {
		if asteroidList[i].name == ast {
			return i
		}
	}
	return -1
}

func getPositionInTime(ast asteroid, time int) (string, int, int) {
	switch time {
	case 1:
		return string(ast.name), ast.time1w, ast.time1h
	case 2:
		return string(ast.name), ast.time2w, ast.time2h
	default:
		return string(ast.name), ast.time3w, ast.time3h
	}
}

func countThird(deltaT1 int, deltaT2 int) {
	for a := range asteroidList {
		ast := asteroidList[a]
		moveWInOne := float64(ast.time2w-ast.time1w) / float64(deltaT1)
		moveHInOne := float64(ast.time2h-ast.time1h) / float64(deltaT1)
		asteroidList[a].time3w += int(math.Floor(moveWInOne*float64(deltaT2))) + ast.time2w
		asteroidList[a].time3h += int(math.Floor(moveHInOne*float64(deltaT2))) + ast.time2h
	}
}

func printField(height int, width int) {
	field := make([][]string, height)
	for h := range field {
		field[h] = make([]string, width)
		for w := range field[h] {
			field[h][w] = "."
		}
	}
	for a := range asteroidList {
		name, w, h := getPositionInTime(asteroidList[a], 0)
		if w < 0 || w > width - 1 || h < 0 || h > height - 1 {
			continue
		}
		if field[h][w] != "." && name > field[h][w] {
			continue
		}
		field[h][w] = name
	}
	for h := range field {
		for w := range field[h] {
			fmt.Printf(field[h][w])
		}
		fmt.Printf("\n")
	}
}

func main() {
	var width, height, time1, time2, time3, deltaT1, deltaT2 int
	_, _ = fmt.Scan(&width, &height, &time1, &time2, &time3)
	deltaT1 = time2 - time1
	deltaT2 = time3 - time2
	for h := 0; h < height; h++ {
		var row1, row2 string
		fmt.Scanf("%s %s", &row1, &row2)
		for w := 0; w < width; w++ {
			if row1[w] != '.' {
				position := checkIfExist(row1[w])
				if position == -1 {
					asteroidList = append(asteroidList, asteroid{name: row1[w], time1w: w, time1h: h})
				} else {
					asteroidList[position].time1w = w
					asteroidList[position].time1h = h
				}
			}
			if row2[w] != '.' {
				position := checkIfExist(row2[w])
				if position == -1 {
					asteroidList = append(asteroidList, asteroid{name: row2[w], time2w: w, time2h: h})
				} else {
					asteroidList[position].time2w = w
					asteroidList[position].time2h = h
				}
			}
		}
	}
	countThird(deltaT1, deltaT2)
	printField(height, width)
}
