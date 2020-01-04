package aoc1

import (
	"reflect"
	"testing"
)

func TestFuelByMass(t *testing.T) {
	type args struct {
		in0 int
		out0 int
	}
	tests := []struct {
		name string
		args args
	}{
		{ "12 should give 2", args{12,2} },
		{ "14 should give 2", args{14,2} },
		{ "1969 should give 654", args{1969,654} },
		{ "100756 should give 33583", args{100756,33583} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFuelValue := FuelByMass(tt.args.in0); !reflect.DeepEqual(gotFuelValue,tt.args.out0) {
				t.Errorf("FuelByMass() = %v , want %v", gotFuelValue,tt.args.out0)
			}
		})
	}
}

func TestFuelRequiredByMass(t *testing.T) {
	type args struct {
		in0 int
		out0 int
	}
	tests := []struct {
		name string
		args args
	}{
		{ "12 should give 2", args{12,2} },
		{ "14 should give 2", args{14,2} },
		{ "1969 should give 654", args{1969,966} },
		{ "100756 should give 33583", args{100756,50346} },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFuelValue := FuelRequiredByMass(tt.args.in0); !reflect.DeepEqual(gotFuelValue,tt.args.out0) {
				t.Errorf("FuelRequiredByMass() = %v , want %v", gotFuelValue,tt.args.out0)
			}
		})
	}
}