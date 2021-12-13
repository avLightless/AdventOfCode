package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(puzzle1("input.txt"))
	fmt.Println(puzzle2("input.txt"))
}

func puzzle1(filename string) int64 {
	// Opening the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Input Variables
	const x1 = 0
	const y1 = 1
	const x2 = 2
	const y2 = 3
	var inputRows [][]int // rows -> x1, y1, x2, y2
	// TODO make dynamic
	var oceanFloor [990][990]int
	biggestX := 0
	biggestY := 0

	// Reading the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txtSplitArrow := strings.Split(scanner.Text(), " -> ")
		split1 := strings.Split(txtSplitArrow[0], ",")
		split2 := strings.Split(txtSplitArrow[1], ",")
		var appendRow []int
		v1, _ := strconv.Atoi(split1[0])
		v2, _ := strconv.Atoi(split1[1])
		v3, _ := strconv.Atoi(split2[0])
		v4, _ := strconv.Atoi(split2[1])
		appendRow = append(appendRow, v1)
		appendRow = append(appendRow, v2)
		appendRow = append(appendRow, v3)
		appendRow = append(appendRow, v4)
		inputRows = append(inputRows, appendRow)
		if appendRow[x1] > biggestX {
			biggestX = appendRow[x1]
		}
		if appendRow[x2] > biggestX {
			biggestX = appendRow[x2]
		}
		if appendRow[y1] > biggestY {
			biggestY = appendRow[y1]
		}
		if appendRow[y2] > biggestY {
			biggestY = appendRow[y2]
		}
	}

	// Map Ocean Floor
	for _, row := range inputRows {
		rowX1 := row[x1]
		rowX2 := row[x2]
		rowY1 := row[y1]
		rowY2 := row[y2]
		if rowX1 == rowX2 {
			if rowY1 > rowY2 {
				for i := rowY2; i <= rowY1; i++ {
					oceanFloor[i][rowX1]++
				}
			} else {
				for i := rowY1; i <= rowY2; i++ {
					oceanFloor[i][rowX1]++
				}
			}
		} else if rowY1 == rowY2 {
			if rowX1 > rowX2 {
				for i := rowX2; i <= rowX1; i++ {
					oceanFloor[rowY1][i]++
				}
			} else {
				for i := rowX1; i <= rowX2; i++ {
					oceanFloor[rowY1][i]++
				}
			}
		}
	}

	// Check for overlap
	overlapCount := 0
	for _, row := range oceanFloor {
		for _, spot := range row {
			if spot > 1 {
				overlapCount++
			}
		}
	}

	return int64(overlapCount)
}

func puzzle2(filename string) int64 {
	// Opening the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Input Variables
	const x1 = 0
	const y1 = 1
	const x2 = 2
	const y2 = 3
	var inputRows [][]int // rows -> x1, y1, x2, y2
	// TODO make dynamic
	var oceanFloor [990][990]int
	biggestX := 0
	biggestY := 0

	// Reading the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txtSplitArrow := strings.Split(scanner.Text(), " -> ")
		split1 := strings.Split(txtSplitArrow[0], ",")
		split2 := strings.Split(txtSplitArrow[1], ",")
		var appendRow []int
		v1, _ := strconv.Atoi(split1[0])
		v2, _ := strconv.Atoi(split1[1])
		v3, _ := strconv.Atoi(split2[0])
		v4, _ := strconv.Atoi(split2[1])
		appendRow = append(appendRow, v1)
		appendRow = append(appendRow, v2)
		appendRow = append(appendRow, v3)
		appendRow = append(appendRow, v4)
		inputRows = append(inputRows, appendRow)
		if appendRow[x1] > biggestX {
			biggestX = appendRow[x1]
		}
		if appendRow[x2] > biggestX {
			biggestX = appendRow[x2]
		}
		if appendRow[y1] > biggestY {
			biggestY = appendRow[y1]
		}
		if appendRow[y2] > biggestY {
			biggestY = appendRow[y2]
		}
	}

	// Map Ocean Floor
	for _, row := range inputRows {
		rowX1 := row[x1]
		rowX2 := row[x2]
		rowY1 := row[y1]
		rowY2 := row[y2]

		if rowX1 == rowX2 {
			if rowY1 > rowY2 {
				for i := rowY2; i <= rowY1; i++ {
					oceanFloor[i][rowX1]++
				}
			} else {
				for i := rowY1; i <= rowY2; i++ {
					oceanFloor[i][rowX1]++
				}
			}
		} else if rowY1 == rowY2 {
			if rowX1 > rowX2 {
				for i := rowX2; i <= rowX1; i++ {
					oceanFloor[rowY1][i]++
				}
			} else {
				for i := rowX1; i <= rowX2; i++ {
					oceanFloor[rowY1][i]++
				}
			}
		} else {
			xRange := getNumbersBetween(rowX1, rowX2)
			yRange := getNumbersBetween(rowY1, rowY2)

			for i := 0; i <= len(xRange)-1; i++ {
				oceanFloor[yRange[i]][xRange[i]]++
			}
		}
	}

	// Check for overlap
	overlapCount := 0
	for _, row := range oceanFloor {
		for _, spot := range row {
			if spot > 1 {
				overlapCount++
			}
		}
	}

	return int64(overlapCount)
}

func getNumbersBetween(num1 int, num2 int) []int {
	var result []int

	if num1 > num2 {
		for i := num1; i >= num2; i-- {
			result = append(result, i)
		}
	} else {
		for i := num1; i <= num2; i++ {
			result = append(result, i)
		}
	}

	return result
}
