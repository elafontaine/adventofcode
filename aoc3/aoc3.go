package aoc3

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type line struct {
	x_axis [2]int
	y_axis [2]int
}

type Wire struct {
	lines []line
}

const NO_CROSSING = 99999999999999999

// Max returns the larger of x or y.
func Max(multiple_int []int) int {
	if multiple_int[0] < multiple_int[1] {
		return multiple_int[1]
	}
	return multiple_int[0]
}

// Min returns the smaller of x or y.
func Min(multiple_int []int) int {
	if multiple_int[0] > multiple_int[1] {
		return multiple_int[1]
	}
	return multiple_int[0]
}

func FindCrossing(instructions string) int {
	crossing := NO_CROSSING
	split := strings.Split(instructions, "\n")
	wire1 := DrawOneWire(split[0])
	wire2 := DrawOneWire(split[1])
	for _, line1 := range wire1.lines[1:] {
		for _, line2  := range wire2.lines[1:] {
			crossing_maybe := walkLinesThrough(line1, line2)
			if crossing_maybe < crossing {
				crossing = crossing_maybe
			}
		}
	}
	return crossing
}

func FindCrossing2(instructions string) int {
	split := strings.Split(instructions, "\n")
	wire1 := DrawOneWire(split[0])
	wire2 := DrawOneWire(split[1])
	steps := NO_CROSSING
	sum1 := 0
	for _, line1 := range wire1.lines[0:] {
		sum2 := 0
		for _, line2  := range wire2.lines[0:] {
			crossing_maybe1, crossing_maybe2 := walkLinesThrough2(line1, line2)
			if crossing_maybe1 != NO_CROSSING && sum1 + sum2 +crossing_maybe1 + crossing_maybe2 < steps {
				steps = sum1 + sum2 + crossing_maybe1 + crossing_maybe2
			}
			sum2 += getLength(line2)
		}
		sum1 += getLength(line1)
	}
	return steps
}

func getLength(line_obj line) int {
	line1_x_diff := line_obj.x_axis[1] - line_obj.x_axis[0]
	line1_y_diff := line_obj.y_axis[1] - line_obj.y_axis[0]
	return int(math.Abs(float64(line1_x_diff))) + int(math.Abs(float64(line1_y_diff)))
}

func walkLinesThrough(line1 line, line2 line) int {
	line1_y_diff := line1.y_axis[1] - line1.y_axis[0]
	line2_y_diff := line2.y_axis[1] - line2.y_axis[0]
	line1_x_diff := line1.x_axis[1] - line1.x_axis[0]
	line2_x_diff := line2.x_axis[1] - line2.x_axis[0]
	if line2_y_diff == 0 && line1_y_diff == 0 {
		return NO_CROSSING
	} else if line2_x_diff == 0 && line1_x_diff == 0 {
		return NO_CROSSING
	} else if line2_y_diff == 0 {
		if Min(line1.y_axis[:]) < line2.y_axis[0] && line2.y_axis[0] < Max(line1.y_axis[:]) {
			for point:= line2.x_axis[0]; point != line2.x_axis[1] ; point+=int(math.Copysign(1,float64(line2_x_diff))){
				if point == line1.x_axis[0] {
					return int(math.Abs(float64(line1.x_axis[0]))) + int(math.Abs(float64(line2.y_axis[0])))
				}
			}
		}
	} else if line2_x_diff == 0 {
		if Min(line1.x_axis[:]) < line2.x_axis[0] && line2.x_axis[0] < Max(line1.x_axis[:]) {
			for point := line2.y_axis[0]; point != line2.y_axis[1]; point += int(math.Copysign(1, float64(line2_y_diff))) {
				if point == line1.y_axis[0] {
					return int(math.Abs(float64(line1.y_axis[0]))) + int(math.Abs(float64(line2.x_axis[0])))
				}
			}
		}
	}
	return NO_CROSSING
}

func walkLinesThrough2(line1 line, line2 line) (int, int) {
	line1_y_diff := line1.y_axis[1] - line1.y_axis[0]
	line2_y_diff := line2.y_axis[1] - line2.y_axis[0]
	line1_x_diff := line1.x_axis[1] - line1.x_axis[0]
	line2_x_diff := line2.x_axis[1] - line2.x_axis[0]
	return_value_y:=NO_CROSSING
	return_value_x:=NO_CROSSING
	if line2_y_diff == 0 && line1_y_diff == 0 {
		return NO_CROSSING, NO_CROSSING
	} else if line2_x_diff == 0 && line1_x_diff == 0 {
		return NO_CROSSING, NO_CROSSING
	} else if line2_y_diff == 0 {
		if Min(line1.y_axis[:]) < line2.y_axis[0] && line2.y_axis[0] < Max(line1.y_axis[:]) {
			x := 0
			for point:= line2.x_axis[0]; point != line2.x_axis[1] ; point+=int(math.Copysign(1,float64(line2_x_diff))){

				if point == line1.x_axis[0] {
					y := 0
					for point2:= line1.y_axis[0]; point2 != line1.y_axis[1] ; point2+=int(math.Copysign(1,float64(line1_y_diff))){

						if point2 == line2.y_axis[0] {
							return_value_y = y
							return x,y
						}
						y++
					}
				}
				x++
			}
		}
	} else if line2_x_diff == 0 {
		if Min(line1.x_axis[:]) < line2.x_axis[0] && line2.x_axis[0] < Max(line1.x_axis[:]) {
			y:=0
			for point := line2.y_axis[0]; point != line2.y_axis[1]; point += int(math.Copysign(1, float64(line2_y_diff))) {

				if point == line1.y_axis[0] {
					x:=0
					for point := line1.x_axis[0]; point != line1.x_axis[1]; point += int(math.Copysign(1, float64(line1_x_diff))) {

						if point == line2.x_axis[0] {
							return_value_x =x
							return x, y
						}
						x++
					}

				}
				y++
			}
		}
	}
	return return_value_x, return_value_y
}
func DrawOneWire(wire_instructions string) Wire {
	wire := Wire{}
	for _, instruction := range strings.Split(wire_instructions,",") {
		wire = addlinetowire(wire,interpret_wire_instruction(string(instruction)))
	}
	return wire
}

func addlinetowire(wire Wire, lineToAdd line) Wire {
	new_wire := Wire{lines:nil}
	if wire.lines == nil{
		return Wire{[]line{lineToAdd}}
	}

	last_line := wire.lines[len(wire.lines)-1]
	new_wire.lines = append(wire.lines, lineadd(last_line,lineToAdd))
	return new_wire
}

func lineadd(official_line line, lineToAdd line) line {
	return line{
		x_axis: [2]int{official_line.x_axis[1] + lineToAdd.x_axis[0],
			official_line.x_axis[1] + lineToAdd.x_axis[1]},
		y_axis: [2]int{official_line.y_axis[1] + lineToAdd.y_axis[0],
			official_line.y_axis[1] + lineToAdd.y_axis[1]},
	}
}
func interpret_wire_instruction(instruction string) line {
	direction := string(instruction[0])
	length, err := strconv.Atoi(string(instruction[1:]))
	if err != nil {
		log.Fatal("conversion error")
	}

	if direction == "U" {
		return line{[2]int{0,0},[2]int{0,length}}
	} else if direction == "D" {
		return line{[2]int{0,0},[2]int{0,-length}}
	} else if direction == "R" {
		return line{[2]int{0,length}, [2]int{0,0},}
	} else if direction == "L" {
		return line{[2]int{0,-length}, [2]int{0,0}}
	}
	return line{
		x_axis: [2]int{0,0},
		y_axis: [2]int{0,0},
	}
}
// keep only the array of joint with their position
