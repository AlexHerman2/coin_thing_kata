package main

import "testing"

func Test_makeChange(t *testing.T) {
	type args struct {
		cents int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "4 cents",
			args: args{
				cents: 4,
			},
			want: " 4P",
		},
		{
			name: "5 cents",
			args: args{
				cents: 5,
			},
			want: " 1N",
		},
		{
			name: "20 cents",
			args: args{
				cents: 20,
			},
			want: " 2D",
		},
		{
			name: "100 cents",
			args: args{
				cents: 100,
			},
			want: " 4Q",
		},
		{
			name: "6 cents",
			args: args{
				cents: 6,
			},
			want: " 1N 1P",
		},
		{
			name: "11 cents",
			args: args{
				cents: 11,
			},
			want: " 1D 1P",
		},
		{
			name: "26 cents",
			args: args{
				cents: 26,
			},
			want: " 1Q 1P",
		},
		{
			name: "16 cents",
			args: args{
				cents: 16,
			},
			want: " 1D 1N 1P",
		},
		{
			name: "41 cents",
			args: args{
				cents: 41,
			},
			want: " 1Q 1D 1N 1P",
		},
		{
			name: "99 cents",
			args: args{
				cents: 99,
			},
			want: " 3Q 2D 4P",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeChange(tt.args.cents); got != tt.want {
				t.Errorf("makeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
