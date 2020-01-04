package aoc6

import (
	"testing"
)

func TestItemGraph_Print(t *testing.T) {
	type fields struct {
		nodes []*Node
		edges map[Node][]*Node
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Empty graph should print empty", fields{
			nodes: nil,
			edges: nil,
		},""},
		{"Graph with one node shouldn't print anything", fields{
			nodes: []*Node{{value: "COM"}},
			edges: nil,
		}, ""},
		{"Graph with one edge should print", fields{
			nodes: []*Node{{value: "COM"},{value:"A"}},
			edges: map[Node][]*Node{Node{value:"COM"}: {&Node{value: "A"}}},
		}, "COM -> A \n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			graph := &ItemGraph{
				nodes: tt.fields.nodes,
				edges: tt.fields.edges,
			}
			if got := graph.Print(); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLinksNumber(t *testing.T) {
	type args struct {
		FileToLoad string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Should load test_file correctly",args{FileToLoad:"test_file"},42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLinksNumber(tt.args.FileToLoad); got != tt.want {
				t.Errorf("GetLinksNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDistanceBetween(t *testing.T) {
	type args struct {
		FileToLoad string
		start      string
		end        string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Distance between you and san", args{
			FileToLoad: "testFile2",
			start:      "YOU",
			end:        "SAN",
		}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDistanceBetween(tt.args.FileToLoad, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("FindDistanceBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}