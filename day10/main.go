package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	mask    int64 
	presses int   
}

func solveMachine(line string) int {
	var targetMask int64
	var buttons []int64

	startBracket := strings.Index(line, "[")
	endBracket := strings.Index(line, "]")
	indicatorStr := line[startBracket + 1 : endBracket]

	for i, char := range indicatorStr {
		if char == '#' {
			targetMask |= (1 << i)
		}
	}

	re := regexp.MustCompile(`\(([\d,]+)\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		parts := strings.Split(match[1], ",")
		var bMask int64
		for _, p := range parts {
			idx, _ := strconv.Atoi(strings.TrimSpace(p))
			bMask |= (1 << idx)
		}
		buttons = append(buttons, bMask)
	}

	// perform BFS
	startState := State{mask: 0, presses: 0}
	queue := []State{startState}
	visited := make(map[int64]bool)
	visited[0] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.mask == targetMask {
			return cur.presses
		}

		for _, btnMask := range buttons {
			nextMask := cur.mask ^ btnMask
			if !visited[nextMask] {
				visited[nextMask] = true
				queue = append(queue, State{
					mask: nextMask,
					presses: cur.presses + 1,
				})
			}
		}
	}
	return -1
}

func part1(file *os.File) int {
	totalPresses := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		presses := solveMachine(line)
		totalPresses += presses
	}


	return totalPresses
}


func main() {
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	answer1 := part1(file)
	fmt.Println("Answer 1:", answer1)
}