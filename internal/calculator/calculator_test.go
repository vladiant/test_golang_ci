package calculator_test

import (
	"errors"
	"testing"

	"github.com/vladiant/test_golang_ci/internal/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 1, 2, 3},
		{"negative", -1, -2, -3},
		{"mixed", -1, 2, 1},
		{"zero", 0, 0, 0},
		{"floats", 1.5, 2.5, 4.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := calculator.Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 5, 3, 2},
		{"negative result", 3, 5, -2},
		{"zero", 0, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := calculator.Subtract(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive", 3, 4, 12},
		{"negative", -3, 4, -12},
		{"zero", 5, 0, 0},
		{"floats", 2.5, 4, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := calculator.Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectedErr error
	}{
		{"positive", 10, 2, 5, nil},
		{"negative", -10, 2, -5, nil},
		{"float result", 7, 2, 3.5, nil},
		{"division by zero", 5, 0, 0, calculator.ErrDivisionByZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := calculator.Divide(tt.a, tt.b)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Divide(%v, %v) error = %v; want %v", tt.a, tt.b, err, tt.expectedErr)
			}
			if err == nil && got != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
