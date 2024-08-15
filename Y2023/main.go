package main

import (
	"Y2023/days"
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "day1":
		days.RunDayOne()
	case "day2":
		days.RunDayTwo()
	default:
		fmt.Println("Invalid day")
	}
}
