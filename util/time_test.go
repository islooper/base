package util

import (
	"fmt"
	"testing"
)

func TestXBeforeDay(t *testing.T) {
	type args struct {
		days int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test1", args{days: 1}, ""},
		{"test2", args{days: 2}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XBeforeDay(tt.args.days); got != tt.want {
				t.Errorf("XBeforeDay() = %v, want %v", got, tt.want)
				fmt.Println(got)
			}
		})
	}
}

func TestNowFormat(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NowFormat(tt.args.format); got != tt.want {
				t.Errorf("NowFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXBeforeDay1(t *testing.T) {
	type args struct {
		days int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XBeforeDay(tt.args.days); got != tt.want {
				t.Errorf("XBeforeDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXBeforeDayFormat(t *testing.T) {
	type args struct {
		days   int
		format string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{
			days:   1,
			format: "2006-01-02",
		}, "",
		},

		{"test2", args{
			days:   2,
			format: "2006-01-02",
		}, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XBeforeDayFormat(tt.args.days, tt.args.format); got != tt.want {
				t.Errorf("XBeforeDayFormat() = %v, want %v", got, tt.want)
				fmt.Println(got)
			}
		})
	}
}
