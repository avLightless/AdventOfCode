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

func puzzle1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontalPosition := 0
	depth := 0

	for scanner.Scan() {
		scnStr := scanner.Text()
		sgmts := strings.Fields(scnStr)
		dir := sgmts[0]
		amnt, err := strconv.Atoi(sgmts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch dir {
		case "forward":
			horizontalPosition += amnt
		case "down":
			depth += amnt
		case "up":
			depth -= amnt
		}
	}

	return horizontalPosition * depth
}

func puzzle2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	horizontalPosition := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		scnStr := scanner.Text()
		sgmts := strings.Fields(scnStr)
		dir := sgmts[0]
		amnt, err := strconv.Atoi(sgmts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch dir {
		case "forward":
			horizontalPosition += amnt
			depth += aim * amnt
		case "down":
			aim += amnt
		case "up":
			aim -= amnt
		}
	}

	return horizontalPosition * depth
}
