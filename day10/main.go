package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

var parensRe = regexp.MustCompile(`\(([\d,]+)\)`)
var curlyRe = regexp.MustCompile(`\{(.*?)\}`)

func solvePart2(file *os.File) int {
	totalMinPresses := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { continue }

		// parse given input
		curlyMatch := curlyRe.FindStringSubmatch(line)
		if curlyMatch == nil { continue }
		
		targetStrs := strings.Split(curlyMatch[1], ",")
		var targets []int
		for _, s := range targetStrs {
			val, _ := strconv.Atoi(strings.TrimSpace(s))
			targets = append(targets, val)
		}

		numEqs := len(targets)
		
		eqToButtons := make([][]int, numEqs)
		btnToEqs := make([][]int, 0)
		
		parenMatches := parensRe.FindAllStringSubmatch(line, -1)
		numVars := len(parenMatches)

		for btnIdx, match := range parenMatches {
			parts := strings.Split(match[1], ",")
			var affects []int
			for _, p := range parts {
				eqIdx, _ := strconv.Atoi(strings.TrimSpace(p))
				if eqIdx < numEqs {
					eqToButtons[eqIdx] = append(eqToButtons[eqIdx], btnIdx)
					affects = append(affects, eqIdx)
				}
			}
			btnToEqs = append(btnToEqs, affects)
		}

		minTotal := math.MaxInt32
		
		// store answers
		assignment := make([]int, numVars)
		for i := range assignment { assignment[i] = -1 }
		
		currentEqValues := make([]int, numEqs)
		
		// backtrack
		var solve func(currentCost int)
		solve = func(currentCost int) {
			if currentCost >= minTotal { return }

			// find the best way to start solving the problem
			bestEq := -1
			minUnknowns := math.MaxInt32

			allSatisfied := true
			
			for i := 0; i < numEqs; i++ {
				rem := targets[i] - currentEqValues[i]
				
				// check how many buttons are available
				unknownCount := 0
				
				for _, btnIdx := range eqToButtons[i] {
					if assignment[btnIdx] == -1 {
						unknownCount++
					}
				}

				if unknownCount == 0 {
					if rem != 0 {
						return
					}
				} else {
					allSatisfied = false
					if unknownCount < minUnknowns {
						minUnknowns = unknownCount
						bestEq = i
					}
				}
			}

			if allSatisfied {
				if currentCost < minTotal {
					minTotal = currentCost
				}
				return
			}

			var varsToAssign []int
			for _, btnIdx := range eqToButtons[bestEq] {
				if assignment[btnIdx] == -1 {
					varsToAssign = append(varsToAssign, btnIdx)
				}
			}
			
			pivotVar := varsToAssign[0]

			rem := targets[bestEq] - currentEqValues[bestEq]
			
			if rem < 0 { return }

			start, end := 0, rem
			if minUnknowns == 1 {
				start, end = rem, rem
			}

			for k := start; k <= end; k++ {
				assignment[pivotVar] = k
				
				validMove := true
				for _, eqIdx := range btnToEqs[pivotVar] {
					currentEqValues[eqIdx] += k
					if currentEqValues[eqIdx] > targets[eqIdx] {
						validMove = false
					}
				}

				if validMove {
					solve(currentCost + k)
				}

				for _, eqIdx := range btnToEqs[pivotVar] {
					currentEqValues[eqIdx] -= k
				}
				assignment[pivotVar] = -1
			}
		}

		solve(0)

		if minTotal != math.MaxInt32 {
			totalMinPresses += minTotal
		}
	}

	return totalMinPresses
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	fmt.Println("Answer Part 2:", solvePart2(file))
}