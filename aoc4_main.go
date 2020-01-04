package main

import (
	"adventofcode/aoc4"
)

func main41(startNumber int, endNumber int) int {
	i:=0
	for number := startNumber; number <= endNumber; number++ {
		if aoc4.CheckPasswords(number) {
			i++
		}
	}
	return i
}

/*
func main() {
	fmt.Print(main41(130254,678275))
} */