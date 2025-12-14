package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Position struct {
	r, c int
}

func abs(x int) int {
	if (x < 0) {
		return x
	}
	return x
}

func unique(intSlice []int) []int {
	if len(intSlice) == 0 {
		return nil
	}
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	sort.Ints(list)
	return list
}

func part1(file *os.File) int {
	var redTiles []Position

	maxArea := 0

	scanner := bufio.NewScanner(file)

	// find the max rows and cols, and keep track of the tile positions
	for scanner.Scan() {
		line := scanner.Text()
		position := strings.Split(line, ",")

		c, _ := strconv.Atoi(strings.TrimSpace(position[0]))
		r, _ := strconv.Atoi(strings.TrimSpace(position[1]))

		redTiles = append(redTiles, Position{r, c})
	}

	sort.Slice(redTiles, func(i int, j int) bool {
		return redTiles[i].r < redTiles[j].r
	})

	for i := 0; i < len(redTiles); i++ {
		for j := 0; j < len(redTiles); j++ {
			width := max(1, abs(redTiles[i].r - redTiles[j].r) + 1)
			height := max(1, abs(redTiles[i].c - redTiles[j].c) + 1)

			maxArea = max(maxArea, width * height)
		}
	}

	return maxArea
}

func part2(file *os.File) int {
	var redTiles []Position
	var distinctX, distinctY []int

	maxArea := 0
	ROWS, COLS := 0, 0

	scanner := bufio.NewScanner(file)

	// find the max rows and cols, and keep track of the tile positions
	for scanner.Scan() {
		line := scanner.Text()
		position := strings.Split(line, ",")

		c, _ := strconv.Atoi(strings.TrimSpace(position[0]))
		r, _ := strconv.Atoi(strings.TrimSpace(position[1]))

		redTiles = append(redTiles, Position{r, c})
	}

	distinctX = append(distinctX, 0, 200000)
	distinctY = append(distinctY, 0, 200000)

	for _, p := range redTiles {
		distinctX = append(distinctX, p.c)
		distinctY = append(distinctY, p.r)

		distinctX = append(distinctX, p.c + 1)
		distinctY = append(distinctY, p.r + 1)
	}

	sort.Ints(distinctX)
	sort.Ints(distinctY)
	distinctX = unique(distinctX)
	distinctY = unique(distinctY)

	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for i, x := range distinctX { xMap[x] = i}
	for i, y := range distinctY { yMap[y] = i}

	ROWS, COLS = len(distinctY) - 1, len(distinctX) - 1

	cellWidths := make([]int, COLS)
	cellHeights := make([]int, ROWS)

	for i := 0; i < COLS; i++ { cellWidths[i] = distinctX[i + 1] - distinctX[i]}
	for i := 0; i < ROWS; i++ { cellHeights[i] = distinctY[i + 1] - distinctY[i]}

	matrix := make([][]int, ROWS)

	for i := range matrix {
		matrix[i] = make([]int, COLS)
	}

	n := len(redTiles)

	for i := 0; i < n; i++ {
		p1 := redTiles[i]
		p2 := redTiles[(i + 1) % n]

		c1, c2 := xMap[p1.c], xMap[p2.c]
		r1, r2 := yMap[p1.r], yMap[p2.r]

		if r1 == r2 {
			start, end := c1, c2
			if start > end { start, end = end, start}
			for c := start; c <= end; c++ {
				matrix[r1][c] = 1
			}
		} else if c1 == c2 {
			start, end := r1, r2
			if start > end { start, end = end, start}
			for r := start; r <= end; r++ {
				matrix[r][c1] = 1
			}
		}
	}

	stack := []Position{{0, 0}}
	directions := []Position{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited := make(map[Position]bool)
	visited[Position{0, 0}] = true

	for len(stack) > 0 {
		curr := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		for _, d := range directions {
			dr, dc := curr.r + d.r, curr.c + d.c

			if dr >= 0 && dr < ROWS && dc >= 0 && dc < COLS {
				if !visited[Position{dr, dc}] && matrix[dr][dc] != 1 {
					visited[Position{dr, dc}] = true
					matrix[dr][dc] = 2
					stack = append(stack, Position{dr, dc})
				}
			}
		}
	}

	fmt.Println("Started Calculating for pSum")

	psum := make([][]int, ROWS)
	for r := 0; r < ROWS; r++ {
		psum[r] = make([]int, COLS)
		for c := 0; c < COLS; c++ {
			val := 0
			if matrix[r][c] != 2 {
				val = cellHeights[r] * cellWidths[c]
			}

			top := 0; if r > 0 { top = psum[r-1][c] }
			left := 0; if c > 0 { left = psum[r][c-1] }
			topLeft := 0; if r > 0 && c > 0 { topLeft = psum[r-1][c-1] }

			psum[r][c] = val + top + left - topLeft
		}
	}

	fmt.Println("Started Calculating for MaxArea")

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := redTiles[i]
			p2 := redTiles[j]

			rMin, rMax := p1.r, p2.r
			if rMin > rMax { rMin, rMax = rMax, rMin }
			cMin, cMax := p1.c, p2.c
			if cMin > cMax { cMin, cMax = cMax, cMin }

			width := cMax - cMin + 1
			height := rMax - rMin + 1
			expectedArea := width * height

			idxR2, idxC2 := yMap[rMax], xMap[cMax]
			idxR1, idxC1 := yMap[rMin], xMap[cMin]

			actualArea := psum[idxR2][idxC2]
			if idxR1 > 0 { actualArea -= psum[idxR1 - 1][idxC2] }
			if idxC1 > 0 { actualArea -= psum[idxR2][idxC1 - 1] }
			if idxR1 > 0 && idxC1 > 0 { actualArea += psum[idxR1 - 1][idxC1-1] }

			if actualArea == expectedArea {
				if expectedArea > maxArea {
					maxArea = expectedArea
				}
			}
		}
	}

	return maxArea
}


func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer1 := part1(file)
	fmt.Println("Answer 1:", answer1)

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	answer2 := part2(file)
	fmt.Println("Answer 2:", answer2)
}