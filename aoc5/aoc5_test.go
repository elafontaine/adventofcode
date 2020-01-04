package aoc5

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
			if got, _ := OpCodeProcessor(tt.args.array, nil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("opCodeProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpCodeProcessor(t *testing.T) {
	type args struct {
		array []int
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"99 becomes 99", args{[]int{99},nil},[]int{99}},
		{"3,0,4,0,99 becomes 5,0,4,0,99", args{[]int{3,0,4,0,99},[]int{5}},[]int{5,0,4,0,99}},
		{"3,0,4,0,99 becomes -1,0,4,0,99", args{[]int{3,0,4,0,99},[]int{-1}},[]int{-1,0,4,0,99}},
		{"3,0,99 becomes -1,0,99", args{[]int{3,0,99},[]int{-1}},[]int{-1,0,99}},
		{"1002,4,3,4,33 becomes 1002,4,3,4,99", args{[]int{1002,4,3,4,33},nil},[]int{1002,4,3,4,99}},
		{"1101,100,-1,4,0 is valid and become 1101,100,-1,4,99", args{[]int{1101,100,-1,4,0},nil}, []int{1101,100,-1,4,99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := OpCodeProcessor(tt.args.array, tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("opCodeProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpCodeProcessor78(t *testing.T) {
	type args struct {
		array []int
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"3,9,8,9,10,9,4,9,99,-1,8 should tell if input is 8 (1)", args{[]int{3,9,8,9,10,9,4,9,99,-1,8},[]int{8}},[]int{1}},
		{"3,9,8,9,10,9,4,9,99,-1,8 should tell if input is not 8 (0)", args{[]int{3,9,8,9,10,9,4,9,99,-1,8},[]int{7}},[]int{0}},
		{"3,3,1108,-1,8,3,4,3,99 should tell if input is 8 (1)", args{[]int{3,3,1108,-1,8,3,4,3,99},[]int{8}},[]int{1}},
		{"3,3,1108,-1,8,3,4,3,99 should tell if input is not 8 (0)", args{[]int{3,3,1108,-1,8,3,4,3,99},[]int{7}},[]int{0}},
		{"3,9,7,9,10,9,4,9,99,-1,8 should tell if input is less than 8 (1)", args{[]int{3,9,7,9,10,9,4,9,99,-1,8},[]int{8}},[]int{0}},
		{"3,9,7,9,10,9,4,9,99,-1,8 should tell if input is less than 8 (1)", args{[]int{3,9,7,9,10,9,4,9,99,-1,8},[]int{7}},[]int{1}},
		{"3,9,7,9,10,9,4,9,99,-1,8 should tell if input is less than 8 (1)", args{[]int{3,9,7,9,10,9,4,9,99,-1,8},[]int{9}},[]int{0}},
		{"3,3,1107,-1,8,3,4,3,99 should tell if input is less than 8 (1)", args{[]int{3,3,1107,-1,8,3,4,3,99},[]int{8}},[]int{0}},
		{"3,3,1107,-1,8,3,4,3,99 should tell if input is less than 8 (1)", args{[]int{3,3,1107,-1,8,3,4,3,99},[]int{7}},[]int{1}},
		{"3,3,1107,-1,8,3,4,3,99 should tell if input is less than 8 (1)", args{[]int{3,3,1107,-1,8,3,4,3,99},[]int{9}},[]int{0}},
		/*{"3,0,4,0,99 becomes -1,0,4,0,99", args{[]int{3,0,4,0,99},[]int{-1}},[]int{-1,0,4,0,99}},
		{"3,0,99 becomes -1,0,99", args{[]int{3,0,99},[]int{-1}},[]int{-1,0,99}},
		{"1002,4,3,4,33 becomes 1002,4,3,4,99", args{[]int{1002,4,3,4,33},nil},[]int{1002,4,3,4,99}},
		{"1101,100,-1,4,0 is valid and become 1101,100,-1,4,99", args{[]int{1101,100,-1,4,0},nil}, []int{1101,100,-1,4,99}},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := OpCodeProcessor(tt.args.array, tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("opCodeProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpCodeProcessor56(t *testing.T) {
	type args struct {
		array []int
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9 should tell if input is 0 (0)", args{[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9},[]int{0}},[]int{0}},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9 should tell if input is not 0 (1)", args{[]int{3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9},[]int{7}},[]int{1}},
		/*{"3,0,4,0,99 becomes -1,0,4,0,99", args{[]int{3,0,4,0,99},[]int{-1}},[]int{-1,0,4,0,99}},
		{"3,0,99 becomes -1,0,99", args{[]int{3,0,99},[]int{-1}},[]int{-1,0,99}},
		{"1002,4,3,4,33 becomes 1002,4,3,4,99", args{[]int{1002,4,3,4,33},nil},[]int{1002,4,3,4,99}},
		{"1101,100,-1,4,0 is valid and become 1101,100,-1,4,99", args{[]int{1101,100,-1,4,0},nil}, []int{1101,100,-1,4,99}},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := OpCodeProcessor(tt.args.array, tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("opCodeProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}