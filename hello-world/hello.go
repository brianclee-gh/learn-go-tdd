package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if len(name) != 0 {
		return englishHelloPrefix + name
	} else {
		return "Hello, world"
	}
	
}

func main() {
	fmt.Printf(Hello("world"))
}