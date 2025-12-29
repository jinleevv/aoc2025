package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func dfs(start string, paths map[string][]string, visited map[string]bool) int {
	numPath := 0

	if start == "out" {
		return 1
	}

	visited[start] = true

	if neighbours, ok := paths[start]; ok {
		for _, neigh := range neighbours {
			if !visited[neigh] {
				numPath += dfs(neigh, paths, visited)
			}
		}
	}

	visited[start] = false

	return numPath
}

func solvePart1(file *os.File) int {
	// copy the input to a map structure
	paths := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		if len(parts) < 2 {
			continue
		}

		key := parts[0]
		value := strings.Fields(parts[1])

		paths[key] = value
	}

	// perform dfs to find the number of paths
	visited := make(map[string]bool)

	return dfs("you", paths, visited)
}


var strToID map[string]int
var nextID int

func getID(s string) int {
	if id, exists := strToID[s]; exists {
		return id
	}
	id := nextID
	nextID++
	strToID[s] = id
	return id
}

var memo [][]int

const (
	MaskNone = 0
	MaskDAC  = 1
	MaskFFT  = 2
	MaskBoth = 3
)

func countPaths(u int, mask int, graph [][]int, outID, dacID, fftID int) int {
	if val := memo[u][mask]; val != -1 {
		return val
	}

	if u == outID {
		if mask == MaskBoth {
			return 1
		}
		return 0
	}

	totalPaths := 0
	for _, v := range graph[u] {
		newMask := mask
		if v == dacID { newMask |= MaskDAC }
		if v == fftID { newMask |= MaskFFT }

		totalPaths += countPaths(v, newMask, graph, outID, dacID, fftID)
	}

	memo[u][mask] = totalPaths
	return totalPaths
}

func solvePart2(file *os.File) int {
	strToID = make(map[string]int)
	nextID = 0
	
	tempPaths := make(map[int][]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 2 { continue }

		u := getID(strings.TrimSpace(parts[0]))
		vals := strings.Fields(parts[1])
		
		for _, vStr := range vals {
			v := getID(vStr)
			tempPaths[u] = append(tempPaths[u], v)
		}
	}

	numNodes := nextID
	graph := make([][]int, numNodes)
	for u, neighbors := range tempPaths {
		graph[u] = neighbors
	}

	if _, ok := strToID["svr"]; !ok { return 0 }
	
	svrID := getID("svr")
	outID := getID("out")
	dacID := getID("dac")
	fftID := getID("fft")

	memo = make([][]int, numNodes)
	for i := range memo {
		memo[i] = []int{-1, -1, -1, -1}
	}

	startMask := 0
	if svrID == dacID { startMask |= MaskDAC }
	if svrID == fftID { startMask |= MaskFFT }

	return countPaths(svrID, startMask, graph, outID, dacID, fftID)
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil { 
		log.Fatal(err) 
	}
	defer file.Close()

	fmt.Println("Answer 2:", solvePart2(file))
}