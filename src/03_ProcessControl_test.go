package exercises

import (
	"testing"
)

func TestSwitch(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"Apple", "Apple match", false},
		{"Google", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Switch(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Switch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Switch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIf(t *testing.T) {
	type args struct {
		age int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Under 18", args{12}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.age); got != tt.want {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}
