package pairsum

import "testing"

func TestHasPairWithSum2(t *testing.T) {
	type args struct {
		nums []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test has sum",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				sum:  17,
			},
			want: true,
		},
		{
			name: "test has no pair",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				sum:  18,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasPairWithSum2(tt.args.nums, tt.args.sum); got != tt.want {
				t.Errorf("HasPairWithSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
