package snippet

import (
	"testing"
)

func Test_get_hamming_weight(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "test_3", args: args{number: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := get_hamming_weight(tt.args.number); got != tt.want {
				t.Errorf("get_hamming_weight() = %v, want %v", got, tt.want)
			}
		})
	}
}
