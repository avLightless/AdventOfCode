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
	var (
		amountOfZero = [12]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		amountOfOne  = [12]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		gamma        = ""
		epsylon      = ""
	)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		if err != nil {
			log.Fatal(err)
		}
		for i, v := range row {
			val, _ := strconv.Atoi(v)
			if val == 0 {
				amountOfZero[i]++
			} else {
				amountOfOne[i]++
			}
		}
	}

	for i, v := range amountOfOne {
		if v > amountOfZero[i] {
			gamma += "1"
			epsylon += "0"
		} else {
			gamma += "0"
			epsylon += "1"
		}
	}

	gammaDec, err := strconv.ParseInt(gamma, 2, 64)
	epsylonDec, err := strconv.ParseInt(epsylon, 2, 64)

	return gammaDec * epsylonDec
}

func puzzle2(filename string) int64 {
	var stringArray [][]string

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		if err != nil {
			log.Fatal(err)
		}
		stringArray = append(stringArray, row)
	}

	mez := jajne(stringArray, 0)
	csupor := strings.Join(mez[0], "")
	malacka, err := strconv.ParseInt(csupor, 2, 64)

	mez2 := jajne2(stringArray, 0)
	csupor2 := strings.Join(mez2[0], "")
	malacka2, err := strconv.ParseInt(csupor2, 2, 64)

	return malacka * malacka2
}

func micimacko(stringArray [][]string, i int, c1 int, c0 int) (c1r int, c0r int) {
	for _, v := range stringArray {
		//fmt.Println(v)
		//fmt.Println(v[0])
		if v[i] == "1" {
			c1++
			fmt.Println("inside function c1 is: ")
			fmt.Println(c1)
		} else {
			c0++
			fmt.Println("inside function c0 is: ")
			fmt.Println(c0)
		}
	}
	return c1, c0
}

func jajne(stringArray [][]string, index int) (result [][]string) {
	fmt.Println("array length")
	fmt.Println(len(stringArray))
	if len(stringArray) == 1 {
		return stringArray
	}
	i := index

	c1 := 0
	c0 := 0
	fmt.Println("c1 BEFORE = ")
	fmt.Println(c1)
	fmt.Println("c0 BEFORE = ")
	fmt.Println(c0)
	c1, c0 = micimacko(stringArray, i, c1, c0)
	fmt.Println("c1 AFTER = ")
	fmt.Println(c1)
	fmt.Println("c0 AFTER = ")
	fmt.Println(c0)

	var newArray [][]string

	if c1 >= c0 {
		fmt.Println("C1 IS LARGER FGGGGGGGGGGGGGGGGGGGGGGGGGGG")
		fmt.Println(len(stringArray))
		for _, v := range stringArray {
			if v[i] == "1" {
				newArray = append(newArray, v)
			}
		}
		fmt.Println(len(newArray))
	} else {
		fmt.Println("C0 IS LARGER FGGGGGGGGGGGGGGGGGGGGGGGGGGG")
		fmt.Println(len(stringArray))
		for _, v2 := range stringArray {
			if v2[i] == "0" {
				newArray = append(newArray, v2)
			}
		}
	}
	result = jajne(newArray, index+1)

	return result
}

func micimacko2(stringArray [][]string, i int, c1 int, c0 int) (c1r int, c0r int) {
	for _, v := range stringArray {
		if v[i] == "1" {
			c1++
		} else {
			c0++
		}
	}
	return c1, c0
}

func jajne2(stringArray [][]string, index int) (result [][]string) {
	if len(stringArray) == 1 {
		return stringArray
	}

	i := index
	c1 := 0
	c0 := 0

	c1, c0 = micimacko2(stringArray, i, c1, c0)

	var newArray [][]string

	if c1 < c0 {
		for _, v := range stringArray {
			if v[i] == "1" {
				newArray = append(newArray, v)
			}
		}
	} else {
		for _, v := range stringArray {
			if v[i] == "0" {
				newArray = append(newArray, v)
			}
		}
	}
	result = jajne2(newArray, index+1)

	return result
}
