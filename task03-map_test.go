package homework

import (
	"reflect"
	"testing"
)

func Test_sortMapValues(t *testing.T) {
	type args struct {
		input map[int]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name:       "sort",
			args:       args{map[int]string{2: "a", 0: "b", 1: "c"}},
			wantResult: []string{"b", "c", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := sortMapValues(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("sortMapValues() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
