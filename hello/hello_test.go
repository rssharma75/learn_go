package hello_test

import (
	"fmt"
	"testing"

	"github.com/rssharma75/learn_go/hello"
)

func Test(t *testing.T) {

	t.Run("Test Greeting", func(t *testing.T) {
		name := "Pranav"
		want := fmt.Sprintf(hello.HELLO_TEMPLATE, name)
		got := hello.Greet(name)

		if got != want {
			t.Errorf("Test failed, expected:%v got:%v", want, got)
		}
	})

	t.Run("In Spanish", func(t *testing.T) {
		name := "Jose"
		want := "Hola, " + name + "!"

		got := hello.Hello(name, "spanish")

		if got != want {
			t.Errorf("Test failed, got:%v, want:%v", got, want)
		}

	})

	t.Run("In French", func(t *testing.T) {
		name := "Jean"
		want := "Mercy, " + name + "!"

		got := hello.Hello(name, "french")

		if got != want {
			t.Errorf("Test failed, got:%v, want:%v", got, want)
		}

	})
}
