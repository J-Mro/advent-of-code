package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// invalid ID: any ID which is made only of some sequence of digits repeated twice
// e.g. 55, 6464, 123123 are INVALID
// no numbers have leading zeroes (0101 is invalid, 101 is valid)
// find all invalid ids in given range

// for part 2:
// many odd numbers are not prime therefore cannot assume all odd length ids can only have a repeated pattern of 0000000 (for example)
// e.g. you could have an id length of 15 that could be something like 135135135135135

// Old part 2 logic
func isSameDigitRepeated(idString string, idNum int, idLength int, sumInvalid int) bool {
	numSameDigit := 0
	l := 0
	for range idLength {
		// The repeating string for an invalid id can repeat at least twice
		if idString[l] == idString[0] {
			numSameDigit++
		}
		l++
	}

	if numSameDigit == idLength {
		fmt.Println(idNum)
		sumInvalid += idNum
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sumInvalid := 0
	for scanner.Scan() {
		line := scanner.Text()
		idRanges := strings.Split(line, ",")
		for i := 0; i < len(idRanges); i++ {
			singleRange := idRanges[i]
			rangeLimits := strings.Split(singleRange, "-")
			start, startErr := strconv.Atoi(rangeLimits[0])
			if startErr != nil {
				log.Fatal(startErr)
			}
			end, endErr := strconv.Atoi(rangeLimits[1])
			if endErr != nil {
				log.Fatal(endErr)
			}
			for j := start; j <= end; j++ {
				idString := strconv.Itoa(j)
				idLength := len(idString)

				// checking for repeating patterns in strings
				isRepeated := false
				for i := 1; i < idLength; i++ {
					substring := strings.Repeat(idString[:i], idLength/i)
					if idLength%i == 0 && substring == idString {
						isRepeated = true
					}
				}
				if isRepeated {
					sumInvalid += j
				}

				// // Part 1 logic
				// if idLength%2 == 0 {
				// 	if idString[0:idLength/2] == idString[idLength/2:] {
				// 		sumInvalid += j
				// 	}
				// }
			}
		}
	}
	fmt.Println(sumInvalid)
}
