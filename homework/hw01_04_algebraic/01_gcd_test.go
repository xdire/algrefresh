package hw01_04_algebraic

import "testing"

func TestGCDMod(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Finding GCD with modulo",
			args:args{
				a: 32,
				b: 56,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDMod(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCDMod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCDSub(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Finding GCD with subtraction",
			args:args{
				a: 32,
				b: 56,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDSub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCDSub() = %v, want %v", got, tt.want)
			}
		})
	}
}