package main

import "testing"

func TestOkDivision(t *testing.T) {
	_, err := Divide(10.0, 5.0)
	if err != nil {
		t.Error("There should be no error here")
	}
}

func TestNokDivision(t *testing.T) {
	_, err := Divide(5.0, 0.0)
	if err == nil {
		t.Error("We want error here")
	}
}

var tests = []struct {
	Name     string
	X        float32
	Y        float32
	Expected float32
	IsError  bool
}{
	{"valid", 10.0, 2.0, 5.0, false},
	{"invalid", 4.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		v, err := Divide(tt.X, tt.Y)
		if tt.IsError {
			if err == nil {
				t.Error("Expected error")
			}
		} else {
			if err != nil {
				t.Error("There should be no error here", err.Error())
			}
			if v != tt.Expected {
				t.Errorf("We expected %v but got %v", tt.Expected, v)
			}
		}
	}
}
