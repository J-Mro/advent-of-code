package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	totalMaxNums := 0

	for scanner.Scan() {
		// I want to get one index of this string and add another index

		line := scanner.Text()
		currentCombo := ""
		maxNum := 0
		for i := 0; i < len(line); i++ {
			for j := 0; j < len(line); j++ {
				// batteries can't be rearranged:
				// the first loop index must always "be on the left" of the second loop index
				if i >= j {
					continue
				}
				currentCombo = string(line[i]) + string(line[j])
				parsed, err := strconv.Atoi(currentCombo)
				if err != nil {
					log.Fatal("Error in the conversion", err)
				}
				if parsed > maxNum {
					maxNum = parsed
				}
			}
		}
		totalMaxNums += maxNum
	}
	fmt.Println(totalMaxNums)

}
