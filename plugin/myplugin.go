package main

import (
	"log"
)

type greeting string

func (g greeting) Greet() {
	log.Println("Hello Universe")
}

// exported as symbol named "Greeter"
var Greeter greeting
