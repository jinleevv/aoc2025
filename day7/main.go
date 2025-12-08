package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	r, c int
}

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	initiated := false
	beamPositions := make(map[int]bool)
	res := 0
	for scanner.Scan() {
		line := scanner.Text()


		if initiated {
			for i := 0; i < len(line); i++ {
				// encounter seperator '^'
				if line[i] == '^' && beamPositions[i] {
					res += 1
					delete(beamPositions, i)
					if i - 1 > 0 {
						beamPositions[i - 1] = true
					} 
					if i + 1 < len(line) {
						beamPositions[i + 1] = true
					}
				}
			}

		} else {
			// find initial position
			intialPosition := strings.IndexByte(line, 'S')
			if intialPosition != -1 {
				beamPositions[intialPosition] = true
				initiated = true
			}
		}
		// fmt.Println("Current res:", res)
	}
	return res
}

var visited map[Position]int

func dfs(matrix []string, r int, c int, rows int, cols int) int {
	if r < 0 || c < 0 || r >= rows ||  c >= cols {
		return 0
	}

	if r == rows - 1 {
		return 1
	}
	nextPosition := Position{r, c}
	if val, ok := visited[nextPosition]; ok {
		return val
	}

	res := 0
	char := matrix[r][c]

	if char == '^' {
		left := dfs(matrix, r + 1, c - 1, rows, cols)
		right := dfs(matrix, r + 1, c + 1, rows, cols)
		res = left + right
	} else {
		res = dfs(matrix, r + 1, c, rows, cols)
	}

	visited[nextPosition] = res

	return res
}

func part2(file *os.File) int {
	var matrix []string

	scanner := bufio.NewScanner(file)
	visited = make(map[Position]int)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}
	ROWS := len(matrix)
	COLS := len(matrix[0])
	startCol := 0

	for i := 0; i < COLS; i++ {
		if matrix[0][i] == 'S' {
			startCol = i
			break
		}
	}
	
	res = dfs(matrix, 0, startCol, ROWS, COLS)

	return res
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	answer1 := part1(file)

	_, err = file.Seek(0, 0)
    if err != nil {
        log.Fatal(err)
    }

	answer2 := part2(file)

	fmt.Println("Part 1 answer:", answer1)
	fmt.Println("Part 2 answer:", answer2)
}