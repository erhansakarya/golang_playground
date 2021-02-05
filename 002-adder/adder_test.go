package main

import "testing"

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, sum int, expected int) {
		t.Helper()

		if sum != expected {
			t.Errorf("got %d want %d", sum, expected)
		}
	}

	sum := Add(2, 2)
	expected := 4
	assertCorrectMessage(t, sum, expected)

}
