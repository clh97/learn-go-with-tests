package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("waldo", "spanish")
		want := "hola, waldo!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in english", func(t *testing.T) {
		got := Hello("waldo", "english")
		want := "hello, waldo!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("waldo", "french")
		want := "bonjour, waldo!"

		assertCorrectMessage(t, got, want)
	})

}
