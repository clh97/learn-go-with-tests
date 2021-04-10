package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("waldo")
		want := "hello, waldo!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello, world! when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "hello, world!"

		assertCorrectMessage(t, got, want)
	})

}
