package main

// import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return "Hola, " + name
	}

	if language == "French" {
		return "Bonjour, " + name
	}

	return englishHelloPrefix + name
}

/* func main() {
	fmt.Println("Hello, World!")
} */
