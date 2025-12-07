package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
				if line[i] == '^' && beamPositions[i]{
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

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer := part1(file)
	fmt.Println("Part 1 answer:", answer)
}