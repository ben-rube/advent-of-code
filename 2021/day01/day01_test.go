package main

import "testing"

func TestProcessesData(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "TestIncrCount",
			args: args{fileName: "./day01_test_input.txt"},
			want: 7,
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ProcessesData(tt.args.fileName)
			if got != tt.want {
				t.Errorf("ProcessesData() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ProcessesData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}