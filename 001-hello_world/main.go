package main

import (
	"fmt"
)

const spanish = "Spanish"
const german = "Hallo"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const germanHelloPrefix = "Hallo, "

// SayHello prints "Hello, name" by arguments.
func SayHello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case german:
		prefix = germanHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(SayHello("Erhan", "German"))
}
