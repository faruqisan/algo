package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test success",
			args: args{
				n: 9,
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibonacci(tt.args.n); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintFibonacci(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test success",
			args: args{
				n: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintFibonacci(tt.args.n)
		})
	}
}
