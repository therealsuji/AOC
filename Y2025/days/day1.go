package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func Part1() {
	var fileName = "./testData/day1.txt"
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading file")
	}

	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))
	currentNum := 50
	zeroCount := 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		direction, numStr := string(line[0]), line[1:]
		num, _ := strconv.Atoi(numStr)
		if strings.ToLower(direction) == "l" {
			move := currentNum - num
			currentNum = mod(move, 100)
		} else {
			move := currentNum + num
			currentNum = mod(move, 100)
		}

		if currentNum == 0 {
			zeroCount++
		}
	}
	fmt.Println(zeroCount)
}

func Part2() {
	var fileName = "./testData/day1.txt"
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading file")
	}

	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split((bufio.ScanLines))
	currentNum := 50
	zeroCount := 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		direction, numStr := string(line[0]), line[1:]
		num, _ := strconv.Atoi(numStr)
		for range num {
			if strings.ToLower(direction) == "l" {
				currentNum--
			} else {
				currentNum++
			}
			// reset
			if mod(currentNum, 100) == 0 {
				currentNum = 0
			}
			if currentNum == 0 {
				zeroCount++
			}
		}
	}
	fmt.Println(zeroCount)
}
