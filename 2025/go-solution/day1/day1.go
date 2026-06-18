package main

// dial is circular
// 0 - 99
// dial starts at 50
// password: no. of times the dial is left pointing at 0 after any rotation in the sequence

import (
	"bufio" // File handling
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	dialStart := 50
	numZeroes := 0

	filePath := "input.txt" // Relative path
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err) // error handling conventions in go
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // method to read a file one line at a time
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() { // advances the scanner to the next line
		line := scanner.Text()
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		fmt.Println(line)

		for range distance {
			if direction == 'L' {
				if dialStart == 0 {
					numZeroes++
					dialStart = 100
				}
				dialStart--
			} else {
				dialStart++
				if dialStart == 100 {
					numZeroes++
					dialStart = 0
				}
			}
		}
		if dialStart == 0 {
			numZeroes++
		}

	}
	fmt.Println(numZeroes)
}
