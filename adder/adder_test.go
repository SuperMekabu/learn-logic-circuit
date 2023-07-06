package adder

import "testing"

func TestHalfAdder(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{
			name:  "f, f",
			args:  args{a: false, b: false},
			want:  false,
			want1: false,
		},
		{
			name:  "t, f",
			args:  args{a: true, b: false},
			want:  false,
			want1: true,
		},
		{
			name:  "f, t",
			args:  args{a: false, b: true},
			want:  false,
			want1: true,
		},
		{
			name:  "t, t",
			args:  args{a: true, b: true},
			want:  true,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := HalfAdder(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("HalfAdder() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HalfAdder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFullAdder(t *testing.T) {
	type args struct {
		a bool
		b bool
		c bool
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{
			name: "f, f, f",
			args: args{
				a: false, b: false, c: false,
			},
			want:  false,
			want1: false,
		},
		{
			name: "t, f, f",
			args: args{
				a: true, b: false, c: false,
			},
			want:  false,
			want1: true,
		},
		{
			name: "f, t, f",
			args: args{
				a: false, b: true, c: false,
			},
			want:  false,
			want1: true,
		},
		{
			name: "f, f, t",
			args: args{
				a: false, b: false, c: true,
			},
			want:  false,
			want1: true,
		},
		{
			name: "t, t, f",
			args: args{
				a: true, b: true, c: false,
			},
			want:  true,
			want1: false,
		},
		{
			name: "t, f, t",
			args: args{
				a: true, b: false, c: true,
			},
			want:  true,
			want1: false,
		},
		{
			name: "f, t, t",
			args: args{
				a: false, b: true, c: true,
			},
			want:  true,
			want1: false,
		},
		{
			name: "t, t, t",
			args: args{
				a: true, b: true, c: true,
			},
			want:  true,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FullAdder(tt.args.a, tt.args.b, tt.args.c)
			if got != tt.want {
				t.Errorf("FullAdder() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FullAdder() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
