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

	// Read chosen numbers
	scanner.Scan()
	chosenNumbers := strings.Split(scanner.Text(), ",")

	// Read Bingo Boards
	var bingoBoards [][][]string
	var chosenSpaces [][]int
	var tmpBrd [][]string
	winningNumber, winningBoard := "0", -1

	tmpR := 0
	scanner.Scan()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			bingoBoards = append(bingoBoards, tmpBrd)
			tmpBrd = [][]string{}
			continue
		}
		tmpBrd = append(tmpBrd, strings.Fields(txt))
		tmpR++
	}
	bingoBoards = append(bingoBoards, tmpBrd)
	tmpBrd = [][]string{}

	for _, number := range chosenNumbers {
		if winningBoard != -1 && winningNumber != "0" {
			break
		}
		for bc, board := range bingoBoards {
			if winningBoard != -1 && winningNumber != "0" {
				break
			}
			for rc, row := range board {
				if winningBoard != 0 && winningNumber != "0" {
					break
				}
				for ic, item := range row {
					if winningBoard != 0 && winningNumber != "0" {
						break
					}
					if number == item {
						chosenSpaces = append(chosenSpaces, []int{bc, rc, ic})
					}
				}
			}
			var (
				cC [5]int
				rC [5]int
			)
			for _, space := range chosenSpaces {
				if space[0] != bc {
					continue
				}
				cC[space[2]]++
				rC[space[1]]++
			}
			for _, v := range cC {
				if v == 5 {
					winningNumber = number
					winningBoard = bc
					break
				}
			}

			for _, v := range rC {
				if v == 5 {
					winningNumber = number
					winningBoard = bc
					break
				}
			}
		}
	}

	var chosenBoardNumbers []int
	unMarkedSum := 0

	for _, space := range chosenSpaces {
		if space[0] != winningBoard {
			continue
		}
		tmpNr, _ := strconv.Atoi(bingoBoards[space[0]][space[1]][space[2]])
		chosenBoardNumbers = append(chosenBoardNumbers, tmpNr)
	}

	for ri, row := range bingoBoards[winningBoard] {
		for ci, column := range row {
			isCChosen := false
			for _, space := range chosenSpaces {
				if space[0] == winningBoard && space[1] == ri && space[2] == ci {
					isCChosen = true
				}
			}
			if !isCChosen {
				nr, _ := strconv.Atoi(column)
				unMarkedSum += nr
			}

		}
	}

	tmpWNr, _ := strconv.Atoi(winningNumber)
	return unMarkedSum * tmpWNr
}

func puzzle2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read chosen numbers
	scanner.Scan()
	chosenNumbers := strings.Split(scanner.Text(), ",")

	// Read Bingo Boards
	var bingoBoards [][][]string
	var chosenSpaces [][]int
	var tmpBrd [][]string
	winningNumber := "0"
	var winningBoard []int

	tmpR := 0
	scanner.Scan()
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			bingoBoards = append(bingoBoards, tmpBrd)
			tmpBrd = [][]string{}
			continue
		}
		tmpBrd = append(tmpBrd, strings.Fields(txt))
		tmpR++
	}
	bingoBoards = append(bingoBoards, tmpBrd)
	tmpBrd = [][]string{}

	for _, number := range chosenNumbers {
		//if winningBoard != -1 && winningNumber != "0" {
		//	break
		//}
		for bc, board := range bingoBoards {
			//if winningBoard != -1 && winningNumber != "0" {
			//	break
			//}
			for rc, row := range board {
				//if winningBoard != 0 && winningNumber != "0" {
				//	break
				//}
				for ic, item := range row {
					//if winningBoard != 0 && winningNumber != "0" {
					//	break
					//}
					if number == item {
						chosenSpaces = append(chosenSpaces, []int{bc, rc, ic})
					}
				}
			}
			var (
				cC [5]int
				rC [5]int
			)
			for _, space := range chosenSpaces {
				if space[0] != bc {
					continue
				}
				cC[space[2]]++
				rC[space[1]]++
			}
			for _, v := range cC {
				if v == 5 {
					alreadyInIt := false
					for _, wbv := range winningBoard {
						if wbv == bc {
							alreadyInIt = true
						}
					}
					if !alreadyInIt {
						winningBoard = append(winningBoard, bc)
						winningNumber = number
					}
					//break
				}
			}

			for _, v := range rC {
				if v == 5 {
					alreadyInIt := false
					for _, wbv := range winningBoard {
						if wbv == bc {
							alreadyInIt = true
						}
					}
					if !alreadyInIt {
						winningBoard = append(winningBoard, bc)
						winningNumber = number
					}
					//break
				}
			}
		}
	}

	var chosenBoardNumbers []int
	unMarkedSum := 0

	for i := len(chosenSpaces) - 1; i > 0; i-- {
		brkFlg := false
		if chosenSpaces[i][0] != winningBoard[len(winningBoard)-1] {
			continue
		}
		tmpChosenSpaceNr := bingoBoards[winningBoard[len(winningBoard)-1]][chosenSpaces[i][1]][chosenSpaces[i][2]]
		for i2 := len(chosenNumbers) - 1; i2 > 0; i2-- {

			if chosenNumbers[i2] == winningNumber {
				brkFlg = true
				break
			}
			if chosenNumbers[i2] == tmpChosenSpaceNr {
				// Remove the element at index i from a.
				copy(chosenSpaces[i:], chosenSpaces[i+1:])         // Shift a[i+1:] left one index.
				chosenSpaces[len(chosenSpaces)-1] = []int{0, 0, 0} // Erase last element (write zero value).
				chosenSpaces = chosenSpaces[:len(chosenSpaces)-1]  // Truncate slice.
				break
			}
		}
		if brkFlg {
			//break
		}
	}

	for _, space := range chosenSpaces {
		if len(winningBoard) != 0 {
			if space[0] != winningBoard[len(winningBoard)-1] {
				continue
			}
			tmpNr, _ := strconv.Atoi(bingoBoards[space[0]][space[1]][space[2]])
			chosenBoardNumbers = append(chosenBoardNumbers, tmpNr)
		}
	}

	if len(winningBoard) != 0 {
		for ri, row := range bingoBoards[winningBoard[len(winningBoard)-1]] {
			for ci, column := range row {
				isCChosen := false
				for _, space := range chosenSpaces {
					if space[0] == winningBoard[len(winningBoard)-1] && space[1] == ri && space[2] == ci {
						isCChosen = true
					}
				}
				if !isCChosen {
					nr, _ := strconv.Atoi(column)
					unMarkedSum += nr
				}
			}
		}
	}

	tmpWNr, _ := strconv.Atoi(winningNumber)

	return unMarkedSum * tmpWNr
}
