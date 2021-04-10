package main

import "fmt"

func Hello(person string) string {
	if len(person) <= 0 {
		return "hello, world!"
	}
	return fmt.Sprintf("hello, %s!", person)
}

func main() {
	Hello("michel")
}
