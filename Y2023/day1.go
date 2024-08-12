package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println("Number: ", num)
		total += num
	}

	fmt.Println("Total: ", total)
}
