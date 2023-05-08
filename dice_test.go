package diceimg

import (
	"testing"

	"github.com/tdewolff/canvas"
)

func TestDice_D6(t *testing.T) {

	d := NewBlank(3+0.2, 3+0.2, 0.2, 0.2, canvas.Gray, canvas.Salmon, canvas.Lightgray)

	type args struct {
		face     int
		w        float64
		h        float64
		filename string
		numeric  bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"D6 1 dot", args{1, 3, 3, "d1.svg", false}},
		{"D6 2 dot", args{2, 3, 3, "d2.svg", false}},
		{"D6 3 dot", args{3, 3, 3, "d3.svg", false}},
		{"D6 4 dot", args{4, 3, 3, "d4.svg", false}},
		{"D6 5 dot", args{5, 3, 3, "d5.svg", false}},
		{"D6 6 dot", args{6, 3, 3, "d6.svg", false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d.D6(tt.args.face, tt.args.w, tt.args.h, tt.args.numeric)
			d.WriteOutput(tt.args.filename)
		})
	}
}
