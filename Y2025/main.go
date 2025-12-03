package main

import (
	"Y2025/days"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "day1":
		days.DayOnePart1()
		days.DayOnePart2()
	case "day2":
		days.RunDayTwo()
	default:
		fmt.Println("Invalid day")
	}
}
