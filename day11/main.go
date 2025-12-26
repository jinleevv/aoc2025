package main

import (
	"bufio"
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

// func cacheKey(current string, visited map[string]bool, dac bool, fft bool) string {
// 	keys := make([]string, 0, len(visited))

// 	for key, value := range visited {
// 		if value {
// 			keys = append(keys, key)
// 		}
// 	}

// 	sort.Strings(keys)
// 	visitedStr := strings.Join(keys, ",")

// 	return fmt.Sprintf("%s|%s|%t|%t", current, visitedStr, dac, fft)
// }

// var memo = make(map[string]int)
// func dfs2(start string, paths map[string][]string, visited map[string]bool, dacStatus bool, fftStatus bool) int {
// 	stateKey := cacheKey(start, visited, dacStatus, fftStatus)
// 	if val, ok := memo[stateKey]; ok {
// 		return val
// 	}

// 	numPath := 0

// 	if start == "out" && dacStatus && fftStatus {
// 		return 1
// 	}

// 	if start == "out" {
// 		return 0
// 	}

// 	visited[start] = true

// 	if neighbours, ok := paths[start]; ok {
// 		for _, neigh := range neighbours {
// 			if !visited[neigh] {
// 				newDac := dacStatus || (neigh == "dac")
// 				newFft := fftStatus || (neigh == "fft")
// 				numPath += dfs2(neigh, paths, visited, newDac, newFft)
// 			}
// 		}
// 	}

// 	visited[start] = false
// 	memo[stateKey] = numPath

// 	return numPath
// }


func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// answer1 := solvePart1(file)
	// fmt.Println("Answer 1:", answer1)

	// _, err = file.Seek(0, 0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}