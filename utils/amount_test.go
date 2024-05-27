package utils_test

import (
	"testing"

	"github.com/qiruos/jupiter/utils"
)

func TestAmountToFloat64(t *testing.T) {
	type args struct {
		amount   uint64
		decimals uint8
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1 with decimals 0",
			args: args{
				amount:   1,
				decimals: 0,
			},
			want: 1,
		},
		{
			name: "100 with decimals 3",
			args: args{
				amount:   100000,
				decimals: 3,
			},
			want: 100,
		},
		{
			name: "1000 with decimals 9",
			args: args{
				amount:   1000000000000,
				decimals: 9,
			},
			want: 1000,
		},
		{
			name: "99999.999999999 with decimals 9",
			args: args{
				amount:   99999999999999,
				decimals: 9,
			},
			want: 99999.999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.AmountToFloat64(tt.args.amount, tt.args.decimals); got != tt.want {
				t.Errorf("AmountToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmountToUint64(t *testing.T) {
	type args struct {
		amount   float64
		decimals uint8
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "1 with decimals 0",
			args: args{
				amount:   1,
				decimals: 0,
			},
			want: 1,
		},
		{
			name: "100 with decimals 3",
			args: args{
				amount:   100,
				decimals: 3,
			},
			want: 100000,
		},
		{
			name: "1000 with decimals 9",
			args: args{
				amount:   1000,
				decimals: 9,
			},
			want: 1000000000000,
		},
		{
			name: "99999.999999999 with decimals 9",
			args: args{
				amount:   99999.999999999,
				decimals: 9,
			},
			want: 99999999999999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.AmountToUint64(tt.args.amount, tt.args.decimals); got != tt.want {
				t.Errorf("AmountToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAmountToString(t *testing.T) {
	type args struct {
		amount   uint64
		decimals uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 with decimals 0",
			args: args{
				amount:   1,
				decimals: 0,
			},
			want: "1",
		},
		{
			name: "100 with decimals 3",
			args: args{
				amount:   100000,
				decimals: 3,
			},
			want: "100",
		},
		{
			name: "1000 with decimals 9",
			args: args{
				amount:   1000000000000,
				decimals: 9,
			},
			want: "1000",
		},
		{
			name: "99999.999999999 with decimals 9",
			args: args{
				amount:   99999999999999,
				decimals: 9,
			},
			want: "99999.999999999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.AmountToString(tt.args.amount, tt.args.decimals); got != tt.want {
				t.Errorf("AmountToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
