package utils_test

import (
	"testing"

	"github.com/qiruos/jupiter/utils"
)

func TestTrimStringBetween(t *testing.T) {
	type args struct {
		str   string
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with counter",
			args: args{
				str:   "[16] test string trim",
				start: "[",
				end:   "]",
			},
			want: "test string trim",
		},
		{
			name: "without counter",
			args: args{
				str:   " test string trim",
				start: "[",
				end:   "]",
			},
			want: "test string trim",
		},
		{
			name: "without counter and spaces",
			args: args{
				str:   "test string trim",
				start: "[",
				end:   "]",
			},
			want: "test string trim",
		},
		{
			name: "with empty counter",
			args: args{
				str:   "[] test string trim",
				start: "[",
				end:   "]",
			},
			want: "test string trim",
		},
		{
			name: "with 2 counters",
			args: args{
				str:   "[16] test string [12] trim",
				start: "[",
				end:   "]",
			},
			want: "test string [12] trim",
		},
		{
			name: "with counter in the middle",
			args: args{
				str:   "test string [12] trim",
				start: "[",
				end:   "]",
			},
			want: "test string  trim",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.TrimStringBetween(tt.args.str, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("TrimStringBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
