package main

import (
	"adventofcode/aoc3"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main3() {
	file, err := os.Open("./aoc3/input3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	accumulator := scanner.Text()

	for scanner.Scan() {
		accumulator += scanner.Text() + "\n"
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print(aoc3.FindCrossing(accumulator))
}
func main32() {
	file, err := os.Open("./aoc3/input32")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	accumulator := scanner.Text()

	for scanner.Scan() {
		accumulator += scanner.Text() + "\n"
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Print(aoc3.FindCrossing2(accumulator))
}



/* func main() {
	main32()
} */