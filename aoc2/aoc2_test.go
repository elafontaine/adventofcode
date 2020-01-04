package aoc2

import (
	"reflect"
	"testing"
)

func Test_opCodeProcessor(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"99 becomes 99", args{[]int{99}},[]int{99}},
		{"1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2)", args{[]int{1,0,0,0,99}},[]int{2,0,0,0,99}},
		{"2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6)", args{[]int{2,3,0,3,99}},[]int{2,3,0,6,99}},
		{"2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801)",args{[]int{2,4,4,5,99,0}},[]int{2,4,4,5,99,9801}},
		{"1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99",args{[]int{1,1,1,4,99,5,6,0,99}}, []int{30,1,1,4,2,5,6,0,99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OpCodeProcessor(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("opCodeProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}