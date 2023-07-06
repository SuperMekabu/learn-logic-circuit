package util

import (
	"reflect"
	"testing"
)

func TestBoolBitToDecimal(t *testing.T) {
	type args struct {
		src []bool
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "101",
			args: args{src: []bool{true, false, true}},
			want: 5,
		},
		{
			name: "11000000",
			args: args{src: []bool{true, true, false, false, false, false, false, false}},
			want: 192,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolBitArrayToDecimal(tt.args.src); got != tt.want {
				t.Errorf("BoolBitArrayToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimalToBitAsBool(t *testing.T) {
	type args struct {
		src int32
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "3",
			args: args{src: 3},
			want: []bool{true, true},
		},
		{
			name: "192",
			args: args{src: 192},
			want: []bool{true, true, false, false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecimalToBoolBitArray(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecimalToBoolBitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args[T any] struct {
		src []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[bool]{
		{
			name: "true, false",
			args: args[bool]{src: []bool{true, false}},
			want: []bool{false, true},
		},
		{
			name: "true, false, false",
			args: args[bool]{src: []bool{true, false, false}},
			want: []bool{false, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
