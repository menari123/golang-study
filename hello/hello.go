package main

import (
	"fmt"

	"go-study/greetings"
)

func main() {
	//Get a greeting message and print it
	message := greetings.Hello("Menari")
	fmt.Println(message)
}
