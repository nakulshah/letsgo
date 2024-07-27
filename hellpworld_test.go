package main

import "testing"

func TestSumInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{name: "simple test", args: args{a: 1, b: 2}, want: 3},
		{name: "simple test", args: args{a: 1, b: 2}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumInt8(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SumInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
