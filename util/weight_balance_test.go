package util

import (
	"fmt"
	"testing"
)

func TestWeightRoundRobinBalance_Add(t *testing.T) {
	type args struct {
		name   string
		weight int
	}
	tests := []args{
		{name: "A", weight: 20},
		{name: "B", weight: 20},
		{name: "C", weight: 20},
	}
	wb := new(WeightRoundRobinBalance)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := wb.Add(tt.name, tt.weight); err != nil {
				t.Errorf("Add() error = %v", err)
			}
		})
	}
}

func TestWeightRoundRobinBalance_Next(t *testing.T) {
	type args struct {
		name   string
		weight int
	}
	tests := []args{
		{name: "A", weight: 30},
		{name: "B", weight: 20},
		{name: "C", weight: 20},
	}
	wb := new(WeightRoundRobinBalance)
	for _, tt := range tests {
		if err := wb.Add(tt.name, tt.weight); err != nil {
			t.Errorf("Add() error = %v", err)
		}
	}
	var (
		a1 int
		a2 int
		a3 int
	)
	for i := 0; i < 70; i++ {
		hh := wb.Next()
		fmt.Println(hh)
		switch hh {
		case "A":
			a1++
		case "B":
			a2++
		case "C":
			a3++
		}
	}
	fmt.Println("a1=", a1, " a2=", a2, " a3=", a3)
}
