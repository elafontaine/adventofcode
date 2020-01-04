package aoc4

import "testing"

func TestCheckPasswords(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test is valid password", args{111111}, false},
		{"Test invalid password due to decrease", args{223450}, false},
		{" test no double in number", args{123789}, false},
		{"test they are only double not triplets", args{123444}, false},
		{"Still passes", args{112233}, true},
		{"test multiple groups", args{111122}, true},
		{"test multiple groups", args{122233}, true},
		{"test multiple groups", args{122333}, true},
		{"test multiple groups", args{123334}, false},
		{"test multiple groups", args{223334}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPasswords(tt.args.number); got != tt.want {
				t.Errorf("CheckPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}