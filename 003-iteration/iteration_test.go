package main

import "testing"

func TestRepeat(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, repeat string, expected string) {
		t.Helper()

		if repeat != expected {
			t.Errorf("got %v want %v", repeat, expected)
		}
	}

	repeated := Repeat("a")
	expected := "aaaaa"
	assertCorrectMessage(t, repeated, expected)

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
