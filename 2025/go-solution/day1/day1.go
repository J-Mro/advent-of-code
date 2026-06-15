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
	"strings"
)

func parseInstruction(direction string) int {
	normalised := strings.ToLower(direction)
	sign := ""
	if normalised[0] == 'l' {
		sign += "-"
	} else if normalised[0] == 'r' {
		sign += "+"
	}
	convertedDirection := sign + direction[1:]
	parsed, err := strconv.Atoi(convertedDirection)
	if err != nil {
		log.Fatal("something went wrong:", err)
	}
	return parsed
}

func turnDial(dialPosition, instruction int) int {
	sum := dialPosition + instruction
	return (sum % 100)

}

func main() {
	dialStart := 50
	numZeroes := 0
	filePath := "input.txt" // Relative path
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err) // error handling conventions in go
	}
	defer file.Close()

	fmt.Println("Reading file line by line:")
	scanner := bufio.NewScanner(file) // method to read a file one line at a time
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() { // advances the scanner to the next line
		line := scanner.Text()
		parsed := parseInstruction(line)
		dialStart = turnDial(dialStart, parsed)
		// fmt.Println("current position: ", dialStart)

		if dialStart == 0 {
			numZeroes++
		}
	}
	fmt.Println(numZeroes)

}
