package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Advent of Code | Day 1")

	fmt.Println(puzzle1("input.txt"))
	fmt.Println(puzzle2("input.txt"))
}

func puzzle1(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	amountOfHigher, previousNumber := 0, 0
	isFirstRow := true

	for scanner.Scan() {
		currentNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if !isFirstRow {
			if currentNumber > previousNumber {
				amountOfHigher++
			}
		} else {
			isFirstRow = false
		}

		previousNumber = currentNumber
	}

	return amountOfHigher
}

func puzzle2(filename string) int {
	var intArray []int
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentNumber, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		intArray = append(intArray, currentNumber)
	}

	amountOfHigher, previousBlockSum := 0, 0

	for i, val := range intArray {
		if i+1 == len(intArray)-1 {
			break
		}
		currentBlockSum := val + intArray[i+1] + intArray[i+2]
		if previousBlockSum != 0 {
			if previousBlockSum < currentBlockSum {
				amountOfHigher++
			}
		}
		previousBlockSum = currentBlockSum
	}

	return amountOfHigher
}
