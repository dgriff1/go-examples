package main

import (
	"fmt"
	"testing"
)

func TestParseString(t *testing.T) {
	integer, dec := ParseString("12.24")
	if integer != 12 {
		t.Fatalf("integer is %d", integer)
	}
	if dec != 24 {
		t.Fatalf("dec should be %d", integer)
	}
}

func TestConvertDollars(t *testing.T) {
	dollars := ConvertDollars(1234)
	fmt.Printf("Dollars (%s)", dollars)
}
