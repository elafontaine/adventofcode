package main

import (
	"adventofcode/aoc1"
	"bufio"
	"log"
	"os"
	"strconv"
)

func main1() {
	file, err := os.Open("./aoc2/input2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	accumulator := aoc1.AccumulateFuel()
	for scanner.Scan() {
		module_weigth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fuel := aoc1.FuelRequiredByMass(module_weigth)
		accumulator(fuel)
	}

	print(accumulator(0))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
