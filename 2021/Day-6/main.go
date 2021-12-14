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
	// Opening the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Input Variables
	var fishes []int // fishes is already plural but I need a var for a single fishes
	daysToAge := 80

	// Reading the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strSplit := strings.Split(scanner.Text(), ",")
		for _, value := range strSplit {
			tmpNr, _ := strconv.Atoi(value)
			fishes = append(fishes, tmpNr)
		}
	}

	// Aging the fishes
	for day := 0; day < daysToAge; day++ {
		fishesToAdd := 0
		for fishi, fish := range fishes {
			if fish == 0 {
				fishes[fishi] = 6
				fishesToAdd++
			} else {
				fishes[fishi]--
			}
		}
		for i := 0; i < fishesToAdd; i++ {
			fishes = append(fishes, 8)
		}
	}

	return len(fishes)
}

func puzzle2(filename string) int {
	// Opening the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Input Variables
	var fishes [9]int // fishes is already plural but I need a var for a single fishes
	daysToAge := 256

	// Reading the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strSplit := strings.Split(scanner.Text(), ",")
		for _, value := range strSplit {
			tmpNr, _ := strconv.Atoi(value)
			fishes[tmpNr]++
		}
	}
	fmt.Println(fishes)

	// Aging the fishes
	for day := 0; day < daysToAge; day++ {
		fishesToAdd := 0
		for fishi, fish := range fishes {
			if fishi == 0 {
				fishesToAdd = fish
				fmt.Println("fishes to add: ")
				fmt.Println(fish)
			}
			if fishi != len(fishes)-1 {
				fishes[fishi] = fishes[fishi+1]
			} else {
				fishes[fishi] = 0
			}
		}
		fishes[6] = fishes[6] + fishesToAdd
		fishes[8] = fishes[8] + fishesToAdd
		fmt.Println(fishes)
	}

	fishCount := 0

	for _, fish := range fishes {
		fishCount += fish
	}

	return fishCount
}
