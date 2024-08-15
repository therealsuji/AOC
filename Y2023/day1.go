package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fileName = "./testData/calibrationTexts.txt"
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading file")
	}

	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))

	var total = 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		var length = len(line)
		var number = ""
		var lastNumber = ""
		var potentialTextNumber = ""
		var lastNumberIndex = -1
		var lastPotentialTextNumber = ""
		var lastPotentialTextNumberIndex = -1

		for i := 0; i < length; i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				lastNumber = string(line[i])
				lastNumberIndex = i
				if len(number) == 0 {
					// First number found
					number += lastNumber
				}
			} else {
				potentialTextNumber += string(line[i])
				if num, isNumber := textToNumber(potentialTextNumber); isNumber {
					lastPotentialTextNumberIndex = i
					lastPotentialTextNumber = strconv.Itoa(num)
					potentialTextNumber = ""
					if len(number) == 0 {
						number += strconv.Itoa(num)
					}
				}
			}
			// Once we have reached the end of the line, add the last number to the number
			if i == length-1 {
				if lastNumberIndex > lastPotentialTextNumberIndex {
					number += lastNumber
				}
				if lastPotentialTextNumberIndex > lastNumberIndex {
					number += lastPotentialTextNumber
				}
			}
		}

		num, _ := strconv.Atoi(number)
		fmt.Println("Number: ", num)
		total += num
	}

	fmt.Println("Total: ", total)
}

func textToNumber(ch string) (number int, isNumber bool) {
	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for key, value := range numberMap {
		if strings.Contains(ch, key) {
			return value, true
		}
	}

	return 0, false
}
