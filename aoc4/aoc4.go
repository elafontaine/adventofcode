package aoc4

import "math"

func CheckPasswords(number int) bool {
	if DecreasingNumberCondition(number){
		return false
	}
	if !NoDoubleCondition(number){
		return false
	}
	return true
}

func NoDoubleCondition(number int) bool {
	numbers := []int{}
	countBy := make(map[int]int)
	for i := 0 ; i <= 5 ; i++ {
		numbers = append(numbers,number / int(math.Pow(10, float64(i))) % 10)
	}
	for _, number := range numbers {
		_, exist := countBy[number]
		if exist{
			countBy[number] += 1
		}else {
			countBy[number] = 1
		}
	}
	for _, count := range countBy {
		if count == 2 {
			return true
		}
	}
	return false
}

func DecreasingNumberCondition(number int) bool {
	for i := 6 ; i >= 1 ; i-- {
		if number/int(math.Pow(10, float64(i)))%10 > number/int(math.Pow(10, float64(i-1)))%10 {
			return true
		}
	}
	return false
}



