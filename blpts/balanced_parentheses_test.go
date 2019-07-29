package blpts

import "testing"

func TestIsBalanced(t *testing.T) {
	type args struct {
		parentheses []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test success",
			args: args{
				parentheses: []string{"{", "(", "[", "]", ")", "}"},
			},
			want: true,
		},
		{
			name: "Test success 2",
			args: args{
				parentheses: []string{"{", "(", "[", "]", ")", "}", "{", "}"},
			},
			want: true,
		},
		{
			name: "Test Fail",
			args: args{
				parentheses: []string{"{", "(", "]", "}", "}"},
			},
		},
		{
			name: "Test Fail 2",
			args: args{
				parentheses: []string{"}", "]"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalanced(tt.args.parentheses); got != tt.want {
				t.Errorf("IsBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
