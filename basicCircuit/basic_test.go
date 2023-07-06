package basicCircuit

import (
	"testing"
)

func TestNAND(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true, b:true",
			args: args{
				a: true,
				b: true,
			},
			want: false,
		},
		{
			name: "a:false, b:true",
			args: args{
				a: false,
				b: true,
			},
			want: true,
		},
		{
			name: "a:true, b:false",
			args: args{
				a: true,
				b: false,
			},
			want: true,
		},
		{
			name: "a:false, b:false",
			args: args{
				a: false,
				b: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NAND(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("NAND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNOT(t *testing.T) {
	type args struct {
		a bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true",
			args: args{a: true},
			want: false,
		},
		{
			name: "a:false",
			args: args{a: false},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NOT(tt.args.a); got != tt.want {
				t.Errorf("NOT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAND(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true, b:true",
			args: args{
				a: true,
				b: true,
			},
			want: true,
		},
		{
			name: "a:false, b:true",
			args: args{
				a: false,
				b: true,
			},
			want: false,
		},
		{
			name: "a:true, b:false",
			args: args{
				a: true,
				b: false,
			},
			want: false,
		},
		{
			name: "a:false, b:false",
			args: args{
				a: false,
				b: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AND(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNOR(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true, b:true",
			args: args{
				a: true,
				b: true,
			},
			want: false,
		},
		{
			name: "a:false, b:true",
			args: args{
				a: false,
				b: true,
			},
			want: false,
		},
		{
			name: "a:true, b:false",
			args: args{
				a: true,
				b: false,
			},
			want: false,
		},
		{
			name: "a:false, b:false",
			args: args{
				a: false,
				b: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NOR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("NOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOR(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true, b:true",
			args: args{
				a: true,
				b: true,
			},
			want: true,
		},
		{
			name: "a:false, b:true",
			args: args{
				a: false,
				b: true,
			},
			want: true,
		},
		{
			name: "a:true, b:false",
			args: args{
				a: true,
				b: false,
			},
			want: true,
		},
		{
			name: "a:false, b:false",
			args: args{
				a: false,
				b: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("OR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXOR(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true, b:true",
			args: args{
				a: true,
				b: true,
			},
			want: false,
		},
		{
			name: "a:false, b:true",
			args: args{
				a: false,
				b: true,
			},
			want: true,
		},
		{
			name: "a:true, b:false",
			args: args{
				a: true,
				b: false,
			},
			want: true,
		},
		{
			name: "a:false, b:false",
			args: args{
				a: false,
				b: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XOR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("XOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigOR(t *testing.T) {
	type args struct {
		a bool
		b bool
		c bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "f, f, f",
			args: args{
				a: false, b: false, c: false,
			},
			want: false,
		},
		{
			name: "t, f, f",
			args: args{
				a: true, b: false, c: false,
			},
			want: true,
		},
		{
			name: "f, t, f",
			args: args{
				a: false, b: true, c: false,
			},
			want: true,
		},
		{
			name: "f, f, t",
			args: args{
				a: false, b: false, c: true,
			},
			want: true,
		},
		{
			name: "t, t, f",
			args: args{
				a: true, b: true, c: false,
			},
			want: true,
		},
		{
			name: "t, f, t",
			args: args{
				a: true, b: false, c: true,
			},
			want: true,
		},
		{
			name: "f, t, t",
			args: args{
				a: false, b: true, c: true,
			},
			want: true,
		},
		{
			name: "t, t, t",
			args: args{
				a: true, b: true, c: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigOR(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("BigOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBigAND(t *testing.T) {
	type args struct {
		a bool
		b bool
		c bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "f, f, f",
			args: args{
				a: false, b: false, c: false,
			},
			want: false,
		},
		{
			name: "t, f, f",
			args: args{
				a: true, b: false, c: false,
			},
			want: false,
		},
		{
			name: "f, t, f",
			args: args{
				a: false, b: true, c: false,
			},
			want: false,
		},
		{
			name: "f, f, t",
			args: args{
				a: false, b: false, c: true,
			},
			want: false,
		},
		{
			name: "t, t, f",
			args: args{
				a: true, b: true, c: false,
			},
			want: false,
		},
		{
			name: "t, f, t",
			args: args{
				a: true, b: false, c: true,
			},
			want: false,
		},
		{
			name: "f, t, t",
			args: args{
				a: false, b: true, c: true,
			},
			want: false,
		},
		{
			name: "t, t, t",
			args: args{
				a: true, b: true, c: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigAND(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("BigAND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXNOR(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a:true, b:true",
			args: args{
				a: true,
				b: true,
			},
			want: true,
		},
		{
			name: "a:false, b:true",
			args: args{
				a: false,
				b: true,
			},
			want: false,
		},
		{
			name: "a:true, b:false",
			args: args{
				a: true,
				b: false,
			},
			want: false,
		},
		{
			name: "a:false, b:false",
			args: args{
				a: false,
				b: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XNOR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("XNOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
