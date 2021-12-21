package integers_test

import (
	"testing"

	"github.com/rssharma75/learn_go/integers"
)

func TestIntegers(t *testing.T) {
	t.Run("Test Adding", func(t *testing.T) {
		first := 2
		second := 3

		want := 5
		got := integers.Add(first, second)

		if want != got {
			t.Error("Test failed, got:%n, expected:%n", got, want)
		}
	})
}
