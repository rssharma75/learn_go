package hello

import "fmt"

func Hello() string {
	return fmt.Sprintf("Hello, World!")
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}
