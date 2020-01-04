package aoc2

func OpCodeProcessor(before []int) []int {
	after, current_pos := processOpCode(before,0)

	for after[current_pos] != 99  {
		after, current_pos = processOpCode(after, current_pos)
	}

	return after
}

func processOpCode(array []int, position int) ([]int,int) {
	opcode := array[position + 0]
	if opcode == 99 {
		return array, position
	}
	if opcode == 1 {
		pos1 := array[array[position+1]]
		pos2 := array[array[position+2]]
		answerPos := array[position+3]
		array[answerPos] = pos1 + pos2
	}
	if opcode == 2 {
		pos1 := array[array[position + 1]]
		pos2 := array[array[position + 2]]
		answerPos := array[position+3]
		array[answerPos] = pos1 * pos2
	}
	return array, position + 4
}