package maps_test

import (
	"testing"

	"github.com/rssharma75/learn_go/maps"
)

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Test failed, expected:%v, got:%v", want, got)
	}

}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Test failed, expected:%s got:%s", got, want)
	}

}

func TestSearch(t *testing.T) {
	dict := maps.Dictionary{"test": "This is a test string"}

	t.Run("Regular Search", func(t *testing.T) {

		got, _ := dict.Search("test")

		want := "This is a test string"

		assertStrings(t, got, want)

	})
	t.Run("Unknown Word", func(t *testing.T) {
		_, err := dict.Search("Unknown")

		assertError(t, err, maps.ErrUnknownTerm)
	})
}

func TestAddWords(t *testing.T) {
	dict := maps.Dictionary{}

	t.Run("Regular Add", func(t *testing.T) {
		dict.Add("test", "This is just a test")

		got, err := dict.Search("test")

		want := "This is just a test"
		if err != nil {
			t.Errorf(("Test failed, should have found the term"))
		}

		assertStrings(t, got, want)

	})
	t.Run("Add existing word", func(t *testing.T) {
		dict.Add("test", "This is just a test")

		err := dict.Add("test", "This is just a test")

		assertError(t, err, maps.ErrWordAlreadyExists)

	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	def := "This is just a test"

	d := maps.Dictionary{word: def}
	t.Run("Existing word", func(t *testing.T) {
		newDefinition := "NeW Definition"
		err := d.Update(word, newDefinition)
		assertError(t, err, nil)

		got, _ := d.Search(word)
		want := newDefinition

		assertStrings(t, got, want)
	})

	t.Run("New Word", func(t *testing.T) {
		word := "test1"
		def := "Def 1"

		err := d.Update(word, def)

		assertError(t, err, maps.ErrUnknownTerm)
	})
}
