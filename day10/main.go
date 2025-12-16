package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func part1(file *os.File) int {
	var positions []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(" ", line)

		indicator := parts[0]
		// find the positions for the light indicators
		for i, c := range indicator {
			if c == '#' { positions = append(positions, i) }
		}

		// process the actions that I can make to calculate the min number of steps
	}


	return 0
}


func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	answer1 := part1(file)
	fmt.Println("Answer 1:", answer1)
}