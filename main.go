package main

import (
	"fmt"
	// flag is a package from the standard library for parsing CLI flags
	"flag"
)

// friendly is a variable that is in scope everywhere in the main package
var friendly bool

// port will hold the value of the TCP port we wish to bind to
var port int

// init is a special function like main and is guaranteed to run before main
func init() {
	// This sets up a boolean flag, held in the `friendly` variable so that it's accessible everywhere
	flag.BoolVar(&friendly, "friendly", false, "whether to print out a greeting")

	flag.IntVar(&port, "port", 7777, "the TCP port to bind to")
	flag.Parse()
}

func main() {
	if friendly {
		fmt.Println("Hello, Gophers!")
	} else {
		fmt.Println("Goodbye, Gophers.")
	}

	portString := fmt.Sprintf(":%d", port)
	fmt.Printf("binding to %s\n", portString)
}
