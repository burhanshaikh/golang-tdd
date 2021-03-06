package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying Hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Printing Hello World when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
