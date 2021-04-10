package main

import "fmt"

func Hello(person string, language string) string {
	hello := greetingPrefix(language)

	if len(person) <= 0 {
		return fmt.Sprintf("%s, world!", hello)
	}

	return fmt.Sprintf("%s, %s!", hello, person)
}

func greetingPrefix(language string) (hello string) {
	switch language {
	case "spanish":
		hello = "hola"
	case "english":
		hello = "hello"
	case "french":
		hello = "bonjour"
	default:
		hello = "?"
	}
	return hello
}

func main() {
	Hello("michel", "english")
}
