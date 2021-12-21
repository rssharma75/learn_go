package iteration_test

import (
	"testing"

	"github.com/rssharma75/learn_go/iteration"
)

func assert(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Test failed: expected:%v, got:%v", want, got)
	}

}

func TestIteration(t *testing.T) {

	t.Run("Default Iteration", func(t *testing.T) {
		got := iteration.Repeat("a", 5)
		want := "aaaaa"
		assert(got, want, t)
	})

	t.Run("Iteration count input", func(t *testing.T) {
		got := iteration.Repeat("a", 10)
		want := "aaaaaaaaaa"

		assert(got, want, t)
	})

}
