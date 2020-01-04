package main

import (
	"adventofcode/aoc6"
	"log"
)
func main61() {
	number := aoc6.GetLinksNumber("./aoc6/input6")
	log.Printf("%v",number)
}

func main62() {
	number := aoc6.FindDistanceBetween("./aoc6/input6","YOU","SAN")

	log.Printf("%v",number)
}
func main() {
	main62()
}