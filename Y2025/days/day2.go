package days

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunDayTwo() {
	var fileName = "./testData/day2.txt"
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error reading file")
	}

	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	const delimiter = ','
	fileScanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if i := bytes.IndexByte(data, delimiter); i >= 0 {
			// Found delimiter return text
			return i + 1, data[0:i], nil
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}

		return 0, nil, nil
	})

	invalidSum := 0
	for fileScanner.Scan() {
		token := strings.TrimSpace(fileScanner.Text())
		parts := strings.Split(token, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)

			mid := len(s) / 2

			left := s[:mid]
			right := s[mid:]

			if left == right {
				invalidSum += i
			}
		}
	}
	println(invalidSum)
}
