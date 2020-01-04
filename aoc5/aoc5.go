package aoc5

import (
	"log"
	"math"
)

func OpCodeProcessor(after []int, inputs []int) ([]int, []int) {
	var outputs []int
	current_pos := 0
	remainingInputs:= inputs

	for after[current_pos] != 99  {
		after, current_pos, remainingInputs, outputs = processOpCode(after, current_pos, remainingInputs, outputs)
	}
	log.Printf("%v",outputs)
	return after, outputs
}

func processOpCode(array []int, position int, inputs []int, outputs []int) ([]int, int, []int, []int) {
	opcode := array[position + 0] % 100


	remainingInputs := inputs
	if opcode == 99 {
		return array, position, nil, nil
	} else if opcode == 1 {
		position = opCode1(array, position)
	} else if opcode == 2 {
		position = opCode2(array, position)
	} else if opcode == 3 {
		remainingInputs, position = OpCode3(array, position, inputs)
	} else if opcode == 4 {
		position,outputs = OpCode4(array, position,outputs)
	} else if opcode == 5 {
		position = OpCode5(array, position)
	} else if opcode == 6 {
		position = OpCode6(array, position)
	} else if opcode == 7 {
		position = OpCode7(array, position)
	} else if opcode == 8 {
		position = OpCode8(array, position)
	} else {
		log.Fatalf("Not a known opCode : %v", opcode)
	}
	return array, position, remainingInputs, outputs
}

func extractParameterValues(array []int, position int, argsNumber int) []int {
	return_value := []int{}
	opcode := array[position]
	var Parameter int

	for i:=0; i < argsNumber  ; i++ {
		Interpretation := opcode / int(math.Pow(10,float64(i)) *100)  % 10
		if Interpretation == 1 {
			Parameter = array[position+i+1]
		} else {
			Parameter = array[array[position+i+1]]
		}
		return_value = append(return_value, Parameter)
	}
	return return_value
}

func OpCode5(array []int, position int) int {
	parameters := extractParameterValues(array, position, 2)
	if parameters[0] != 0 {
		position = parameters[1]
		return position
	}
	return position + 3
}

func OpCode6(array []int, position int) int {
	parameters := extractParameterValues(array, position, 2)
	if parameters[0] == 0 {
		position = parameters[1]
		return position
	}
	return position + 3
}

func OpCode7(array []int, position int) int {
	parameters := extractParameterValues(array, position, 3)
	if parameters[0] < parameters[1] {
		array[array[position+3]] = 1
	} else {
		array[array[position+3]] = 0
	}

	return position + 4
}

func OpCode8(array []int, position int) int {
	parameters := extractParameterValues(array, position, 3)
	if parameters[0] == parameters[1] {
		array[array[position+3]] = 1
	} else {
		array[array[position+3]] = 0
	}

	return position + 4
}

func OpCode4(array []int, position int,outputs []int) (int, []int) {
	parameters := extractParameterValues(array, position, 1)
	pos1 := parameters[0]
	outputs = append(outputs, pos1)
	return position + 2, outputs
}

func OpCode3(array []int, position int, inputs []int) ([]int, int) {
	array[array[position+1]] = inputs[0]
	remainingInputs := []int{}
	if len(inputs) > 1 {
		remainingInputs = inputs[1:]
	}
	return remainingInputs, position + 2
}

func opCode1(array []int, position int) int {
	parameters := extractParameterValues( array, position,3)
	pos1 := parameters[0]
	pos2 := parameters[1]
	answerPos := array[position+3]
	array[answerPos] = pos1 + pos2
	return position + 4
}

func opCode2(array []int, position int) int {
	parameters := extractParameterValues( array, position,3)
	pos1 := parameters[0]
	pos2 := parameters[1]
	answerPos := array[position+3]
	array[answerPos] = pos1 * pos2
	return position + 4
}
