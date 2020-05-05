package exercises

import (
	"reflect"
	"testing"
)

func TestCopySlice(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			"Test A,B => Hello World, B",
			args{[]string{"A", "B"}},
			[]string{"A", "B"},
			[]string{"Hello World", "B"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CopySlice(tt.args.strings)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopySlice() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CopySlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
