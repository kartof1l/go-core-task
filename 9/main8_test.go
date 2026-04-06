package main

import (
	"math"
	"testing"
)

func TestChannelPipeline(t *testing.T) {
	unsigned := make(chan uint8)
	floated := make(chan float64)
	results := make(chan float64, 101)

	go func() {
		defer close(unsigned)
		for i := uint8(0); i <= 100; i++ {
			unsigned <- i
		}
	}()

	go func() {
		defer close(floated)
		for i := range unsigned {
			floated <- float64(i) * float64(i) * float64(i)
		}
	}()

	go func() {
		for f := range floated {
			results <- f
		}
		close(results)
	}()

	count := 0
	for r := range results {
		count++
		expected := float64(count-1) * float64(count-1) * float64(count-1)
		if r != expected {
			t.Errorf("expected %.0f, got %.0f", expected, r)
		}
	}
	if count != 101 {
		t.Errorf("expected 101 values, got %d", count)
	}
}

func TestPowValues(t *testing.T) {
	tests := []struct {
		input    uint8
		expected float64
	}{
		{0, 0},
		{1, 1},
		{2, 8},
		{3, 27},
		{4, 64},
		{5, 125},
		{10, 1000},
	}

	for _, tt := range tests {
		result := math.Pow(float64(tt.input), 3)
		if result != tt.expected {
			t.Errorf("pow(%d,3) = %.0f, want %.0f", tt.input, result, tt.expected)
		}
	}
}