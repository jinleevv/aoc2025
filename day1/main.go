package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Solving Day 1")
	num := 50
	res := 0

	// open a file in GoLang
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalf("Input was not able to load properly: %s", err)
	}
	// make sure the file is closed
	defer file.Close()

	// initialize the sanner
	scanner := bufio.NewScanner(file)

	// start 
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		rotationStr := line[1:]
		
		// convert byte to int
		rotation, err := strconv.Atoi(rotationStr)
		if err != nil {
			log.Printf("Failed to parse number: %s", line)
		}
		
		// make sure you match type with `byte`
		if direction == 'R' {
			num += rotation
		} else {
			num -= rotation
		}

		num = num % 100

		if num == 0 {
			res ++
		}
	}
	// check for any errors while scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// print the result
	fmt.Println(res)
}

// answer: 1040