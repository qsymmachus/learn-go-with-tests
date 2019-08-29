package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	switch language {
	case "en":
		return "Hello, " + name
	case "es":
		return "Hola, " + name
	case "fr":
		return "Bonjour, " + name
	default:
		return "Hello, " + name
	}
}

func main() {
	fmt.Println(Hello("John", "en"))
}
