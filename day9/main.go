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