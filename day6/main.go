package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	start, end int
}

func part1(file *os.File) {
	var matrix [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Fields(line)
		matrix = append(matrix, row)
	}
	ROWS := len(matrix)
	COLS := len(matrix[0])

	res := 0
	for i := 0; i < COLS; i++ {
		curSum := 0
		operation := matrix[ROWS - 1][i]
		for j := 0; j < ROWS - 1; j++ {
			curNum, _ := strconv.Atoi(matrix[j][i]) 
			switch operation {
				case "+":
					curSum += curNum 
				case "*":
					if curSum == 0 {
						curSum = curNum
					} else{
						curSum *= curNum
					}
			}
		}
		res += curSum
	}

	fmt.Println("Part 1 answer:", res)
}

func collectDigits(matrix [][]string, row int, col int, digit int, length int, maxLength int) int {
	returnNum := ""

	for i := 0; i < row - 1; i++ {
		// fmt.Println("digit:", digit)
		if len(matrix[i][col]) >= length {
			difference := maxLength - len(matrix[i][col])
			accessDigit := digit - difference
			curNum := matrix[i][col][accessDigit]
			returnNum = returnNum + string(curNum)
		}
	}

	fmt.Println("returnNum:", returnNum)
	if returnNum == "" {
        return 0
    }

	res, _ := strconv.Atoi(returnNum)
	return res
}

// func part2(file *os.File) {
// 	var matrix []string
// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		matrix = append(matrix, line)
// 	}
// 	ROWS := len(matrix)
// 	COLS := len(matrix[0])

// 	res := 0
// 	// you need to collect the digits
// 	for i := 0; i < COLS; i++ {
// 		curSum := 0
// 		operation := matrix[ROWS - 1][i]
// 		// collect the max length of the number 
// 		maxLength := 0
// 		for j := 0; j < ROWS - 1; j++ {
// 			maxLength = max(maxLength, len(matrix[j][i]))
// 		}
// 		// sum/multiply up all the numbers in the column
// 		for k := 0 ; k < maxLength; k++ {
// 			val := collectDigits(matrix, ROWS, i, k, maxLength - k, maxLength) 
// 			switch operation {
// 				case "+":
// 					curSum += val
// 				case "*":
// 					if curSum == 0 {
// 						curSum = val
// 					} else{
// 						curSum *= val
// 					}
// 			}
// 		}
// 		fmt.Println("curSum:", curSum)
// 		res += curSum
// 	}
// 	fmt.Println("Part 2 answer:", res)
// }

func main() {
	file, err := os.Open("testCase.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(file)

	// _, err = file.Seek(0, 0)
    // if err != nil {
    //     log.Fatal(err)
    // }
	
	// part2(file)
}