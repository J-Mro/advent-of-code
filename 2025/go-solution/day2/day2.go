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
				if idLength%2 == 0 {
					if idString[0:idLength/2] == idString[idLength/2:] {
						sumInvalid += j
					}
				}
			}
		}
	}
	fmt.Println(sumInvalid)
}
