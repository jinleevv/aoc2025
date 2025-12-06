package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func repeat_sequence_1(num int, repeat int) bool {
	numStr := strconv.Itoa(num)
	n := len(numStr)
	for i := 1; i < n; i++ {
		pattern := numStr[:i]
		if strings.Repeat(pattern, repeat) == numStr {
			return true
		}
	}
	return false
}

func repeat_sequence_2(num int) bool {
	numStr := strconv.Itoa(num)
	n := len(numStr)
	for i := 1; i < n; i++ {
		if n % i == 0 {
			pattern := numStr[:i]
			if strings.Repeat(pattern, n / i) == numStr {
				return true
			}
		}
	}
	return false
}

func part1(file *os.File) {
	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			var start, end int
			cleanR := strings.TrimSpace(r)
			_, err := fmt.Sscanf(cleanR, "%d-%d", &start, &end)
			if err != nil {
				log.Println("Failed to parse line:", cleanR)
			}
			for i := start; i < end; i++ {
				if repeat_sequence_1(i, 2) {
					res += i
				}
			}
		}
	}
	fmt.Println("Part 1 answer:", res)
}

func part2(file *os.File) {
	res := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			var start, end int
			cleanR := strings.TrimSpace(r)
			_, err := fmt.Sscanf(cleanR, "%d-%d", &start, &end)
			if err != nil {
				log.Println("Failed to parse line:", cleanR)
			}
			for i := start; i < end; i++ {
				if repeat_sequence_2(i) {
					res += i
				}
			}
		}
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
	
	// make sure you start from the beginning
	_, err = file.Seek(0, 0)
    if err != nil {
        log.Fatal(err)
    }

	part2(file)
}