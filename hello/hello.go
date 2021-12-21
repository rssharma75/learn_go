package hello

import "fmt"

const HELLO_TEMPLATE = "Hello, %v!"
const HELLO_TEMP_ES = "Hola, %v!"
const HELLO_TEMP_FR = "Mercy, %v!"

func Greet(name string) string {
	return fmt.Sprintf(HELLO_TEMPLATE, name)
}

func Hello(name string, lang string) string {

	template := HELLO_TEMPLATE
	switch lang {
	case "spanish":
		template = HELLO_TEMP_ES
	case "french":
		template = HELLO_TEMP_FR
	default:
		template = HELLO_TEMPLATE
	}
	return fmt.Sprintf(template, name)
}
