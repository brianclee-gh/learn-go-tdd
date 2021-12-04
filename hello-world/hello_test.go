package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sam", "French")
		want := "Bonjour, Sam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying Hello, world when empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

}