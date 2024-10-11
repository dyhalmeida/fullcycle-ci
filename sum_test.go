package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	expect := 20
	result := sum(10, 10)
	if result != expect {
		t.Errorf("Expected %d, but got %d", expect, result)
	}
}
