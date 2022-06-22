package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(20, 20)

	if total != 40 {
		t.Errorf("Result Sum is invalid: Result: %d. Expect: %d", total, 40)
	}
}