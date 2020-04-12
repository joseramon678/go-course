package main

import (
	"fmt"
	_ "log"
)

// UPERCASE, PUBLIC FUNC
func Message() string {
	return "Hello world"
}

// Return specific message
func MessageImplicit() (msg string) {
	msg = "Hello Go"
	return
}

func main() {
	fmt.Println(MessageImplicit())
}
