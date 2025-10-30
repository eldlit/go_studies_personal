package main

import (
	"fmt"
	"strings"
)

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const helloSuffix = "!"

func main() {
	fmt.Println(Hello("World", "Spanish"))
}

func Hello(name, language string) string {
	language = strings.ToLower(language)

	switch language {
	case "spanish":
		if name == "" {
			return spanishHelloPrefix + "World" + helloSuffix
		}
		return spanishHelloPrefix + name + helloSuffix

	}
	if name == "" {
		return "Hello World!"
	}
	return englishHelloPrefix + name + helloSuffix

}
