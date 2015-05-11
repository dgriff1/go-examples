package main

import (
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
	if dollars != "twelve hundred and thirty four dollars" {
		t.Fatalf("Bad response %s ", dollars)
	}
}

func TestLookupNumber(t *testing.T) {
	val := LookupNumber(5)
	if val != "five" {
		t.Fatalf("Bad response %s ", val)
	}

	val = LookupNumber(25)
	if val != "twenty five" {
		t.Fatalf("Bad response %s ", val)
	}
}
