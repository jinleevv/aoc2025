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

func collectDigits(matrix []string, row, col int) int {
	returnNum := ""
	for r := 0; r < row - 1; r++ {
		if matrix[r][col] != ' ' {
			returnNum += string(matrix[r][col])
		}
	}

	if returnNum == "" {
        return 0
    }

	res, _ := strconv.Atoi(returnNum)
	return res
}

func solveOperation(matrix []string, start int, end int) int {
	ROWS := len(matrix)
	// find operation
	operator := "+"
	for c := start; c <= end; c++ {
		findOperator := matrix[ROWS - 1][c]
		if findOperator == '+' || findOperator == '*' {
			operator = string(findOperator)
			break
		}
	}

	curSum := 0
	for c := end; c >= start; c-- {
		digits := collectDigits(matrix, ROWS, c)
		switch operator {
			case "+":
				curSum += digits 
			case "*":
				if curSum == 0 {
					curSum = digits
				} else{
					curSum *= digits
				}
		}
	}
	return curSum
}

func part2(file *os.File) {
	var matrix []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}
	ROWS := len(matrix)
	COLS := len(matrix[0])

	res := 0

	// seperate the each operations
	operationStart := 0
	var operations []Operation
	for c := 0; c < COLS; c++ {
		seperator := true
		for r := 0; r < ROWS; r++ {
			if matrix[r][c] != ' ' {
				seperator = false
			}
		}

		if seperator && c > operationStart {
			operations = append(operations, Operation{operationStart, c - 1})
			operationStart = c + 1
		}
	}

	if operationStart < COLS {
		operations = append(operations, Operation{operationStart, COLS - 1})
	}

	for i := len(operations) - 1; i >= 0; i-- {
		val := solveOperation(matrix, operations[i].start, operations[i].end)
		res += val
	}
	
	
	fmt.Println("Part 2 answer:", res)
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	part1(file)

	_, err = file.Seek(0, 0)
    if err != nil {
        log.Fatal(err)
    }
	
	part2(file)
}