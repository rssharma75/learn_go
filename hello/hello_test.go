package hello_test

import (
	"fmt"
	"testing"

	"github.com/rssharma75/learn_go/hello"
)

func Test(t *testing.T) {
	t.Run("Static Hello", func(t *testing.T) {
		want := "Hello, World!"
		got := hello.Hello()

		if got != want {
			t.Errorf("Test failed, got:%v , expected:%v", got, want)
		}
	})

	t.Run("Test Greeting", func(t *testing.T) {
		name := "Pranav"
		want := fmt.Sprintf("Hello, %v!", name)
		got := hello.Greet(name)

		if got != want {
			t.Errorf("Test failed, expected:%v got:%v", want, got)
		}
	})
}
