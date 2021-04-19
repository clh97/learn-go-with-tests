package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "123"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "123"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("aaaa")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		value := "this is just a test"

		err := dict.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "abc"
		value := "def"
		dict := Dictionary{key: value}

		err := dict.Add(key, "123")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		newValue := "new definition"
		dictionary := Dictionary{key: value}

		err := dictionary.Update(key, newValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "test definition"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", key)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
