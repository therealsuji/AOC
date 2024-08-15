package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partTwo() {
	var fileName = "./testData/day1.txt"
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
		var potentialTextNumber = ""

		for i := 0; i < length; i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				potentialTextNumber = ""
				number = string(line[i])
				break
			} else {
				potentialTextNumber += string(line[i])
				if num, isNumber := textToNumber(potentialTextNumber); isNumber {
					potentialTextNumber = ""
					number = strconv.Itoa(num)
					break
				}
			}
		}

		for i := length - 1; i >= 0; i-- {
			ch := string(line[i])
			if _, err := strconv.Atoi(ch); err == nil {
				number += ch
				break
			} else {
				potentialTextNumber = ch + potentialTextNumber
				if num, isNumber := textToNumber(potentialTextNumber); isNumber {
					potentialTextNumber = ""
					number += strconv.Itoa(num)
					break
				}
			}
		}

		num, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Error converting number:", err)
			continue
		}
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

func partOne() {
	var fileName = "./testData/day1.txt"
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
		for i := 0; i < length; i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				lastNumber = string(line[i])
				if len(number) == 0 {
					number += lastNumber
				}
			}
			// Once we have reached the end of the line, add the last number to the number
			if i == length-1 {
				number += lastNumber
			}
		}

		num, _ := strconv.Atoi(number)
		total += num
	}

	fmt.Println("Total: ", total)
}

func RunDayOne() {
	partOne()
	partTwo()
}
