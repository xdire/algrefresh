package hw01_04_algebraic

import "testing"

func TestPow2Bin(t *testing.T) {
	type args struct {
		a float64
		b int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test power through binary",
			args:args{
				a: 5,
				b: 9,
			},
			want: 1953125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow2Bin(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Pow2Bin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPow2Mul(t *testing.T) {
	type args struct {
		a float64
		b int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test power through quadratic multiplications",
			args:args{
				a: 5,
				b: 8,
			},
			want: 390625,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pow2Mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Pow2Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPowIter(t *testing.T) {
	type args struct {
		a float64
		b uint32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test power through iteration",
			args:args{
				a: 5,
				b: 8,
			},
			want: 390625,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PowIter(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("PowIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
