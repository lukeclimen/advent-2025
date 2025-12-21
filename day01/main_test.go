package main

import (
	"testing"
)

func TestSpinDialRightNormal(t *testing.T) {
	starting_number := 10
	num_clicks := 40
	msg := SpinDialRight(starting_number, num_clicks)
	if msg != 50 {
		t.Errorf("Expected 50, received %d", msg)
	}
}

func TestSpinDialRightOverOneHundred(t *testing.T) {
	starting_number := 70
	num_clicks := 404
	msg := SpinDialRight(starting_number, num_clicks)
	if msg != 74 {
		t.Errorf("Expected, 74, received %d", msg)
	}
}

func TestSpinDialRightOverCycle(t *testing.T) {
	starting_number := 70
	num_clicks := 440
	msg := SpinDialRight(starting_number, num_clicks)
	if msg != 10 {
		t.Errorf("Expected, 10, received %d", msg)
	}
}

func TestSpinDialLeftNormal(t *testing.T) {
	starting_number := 40
	num_clicks := 10
	msg := SpinDialLeft(starting_number, num_clicks)
	if msg != 30 {
		t.Errorf("Expected, 30, received %d", msg)
	}
}

func TestSpinDialLeftOverOneHundred(t *testing.T) {
	starting_number := 70
	num_clicks := 404
	msg := SpinDialLeft(starting_number, num_clicks)
	if msg != 66 {
		t.Errorf("Expected, 66, received %d", msg)
	}
}

func TestSpinDialLeftOverCycle(t *testing.T) {
	starting_number := 10
	num_clicks := 440
	msg := SpinDialLeft(starting_number, num_clicks)
	if msg != 70 {
		t.Errorf("Expected, 70, received %d", msg)
	}
}