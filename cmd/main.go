// Package main is the CLI entry point for the calc tool.
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/vladiant/test_golang_ci/internal/calculator"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <operation> <a> <b>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Operations: add, subtract, multiply, divide\n")
		os.Exit(1)
	}

	op := os.Args[1]

	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalf("invalid number %q: %v", os.Args[2], err)
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatalf("invalid number %q: %v", os.Args[3], err)
	}

	var result float64

	switch op {
	case "add":
		result = calculator.Add(a, b)
	case "subtract":
		result = calculator.Subtract(a, b)
	case "multiply":
		result = calculator.Multiply(a, b)
	case "divide":
		result, err = calculator.Divide(a, b)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	default:
		log.Fatalf("unknown operation %q; use add, subtract, multiply, or divide", op)
	}

	fmt.Printf("%.10g\n", result)
}
