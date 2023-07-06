package calculator

import (
	"learn-logics/util"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
	type args struct {
		a []bool
		b []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "1 + 1",
			args: args{
				a: util.DecimalToBoolBitArray(1),
				b: util.DecimalToBoolBitArray(1),
			},
			want: util.DecimalToBoolBitArray(1 + 1),
		},
		{
			name: "3 + 7",
			args: args{
				a: util.DecimalToBoolBitArray(3),
				b: util.DecimalToBoolBitArray(7),
			},
			want: util.DecimalToBoolBitArray(3 + 7),
		},
		{
			name: "15 + 15",
			args: args{
				a: util.DecimalToBoolBitArray(15),
				b: util.DecimalToBoolBitArray(15),
			},
			want: util.DecimalToBoolBitArray(15 + 15),
		},
		{
			name: "16 + 16",
			args: args{
				a: util.DecimalToBoolBitArray(16),
				b: util.DecimalToBoolBitArray(16),
			},
			want: util.DecimalToBoolBitArray(16 + 16),
		},
		{
			name: "195728 + 312931741",
			args: args{
				a: util.DecimalToBoolBitArray(195728),
				b: util.DecimalToBoolBitArray(31293129318),
			},
			want: util.DecimalToBoolBitArray(195728 + 31293129318),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adder(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adder() = %v, want %v", got, tt.want)
			}
		})
	}
}
