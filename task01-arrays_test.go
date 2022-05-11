package homework

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		input [15]float32
	}
	tests := []struct {
		name       string
		args       args
		wantResult float32
	}{
		{
			name:       "one values",
			args:       args{[15]float32{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}},
			wantResult: 1.0},
		{
			name:       "two-zero values",
			args:       args{[15]float32{15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0, 15.0}},
			wantResult: 15.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := average(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("average() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
