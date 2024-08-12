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
		for i := 0; i < length; i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				number += string(line[i])
				break
			}
		}

		for i := length - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				number += string(line[i])
				break
			}
		}

		num, _ := strconv.Atoi(number)
		fmt.Println("Number: ", num)
		total += num
	}

	fmt.Println("Total: ", total)

	fmt.Println("Done")
}
