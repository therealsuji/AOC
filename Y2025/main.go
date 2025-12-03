package main

import (
	"Y2025/days"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "day1":
		// days.Part1()
		days.Part2()
	default:
		fmt.Println("Invalid day")
	}
}
