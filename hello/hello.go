package hello

import "fmt"

const HELLO_TEMPLATE = "Hello, %v!"

func Hello() string {
	return fmt.Sprintf("Hello, World!")
}

func Greet(name string) string {
	return fmt.Sprintf(HELLO_TEMPLATE, name)
}
