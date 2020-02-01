package main

import "testing"

func Test_getShootingPercentage(t *testing.T) {
	type args struct {
		goals    int
		attempts int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getShootingPercentage(tt.args.goals, tt.args.attempts); got != tt.want {
				t.Errorf("getShootingPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
