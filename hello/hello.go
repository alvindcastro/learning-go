package hello

import "fmt"

const spanish = "Spanish"
const french = "French"
const filipino = "Filipino"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const filipinoHelloPrefix = "Hoy, "

// Hello : Initial function
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {

	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case filipino:
		prefix = filipinoHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func hello() {
	fmt.Println(Hello("world", "English"))
}
