package main

import (
	"fmt"
	"rsc.io/quote"
	"go-playground/greetings"
)


func main() {
	fmt.Println(quote.Go())
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
