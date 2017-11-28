package main

import (
	"fmt"
	// flag is a package from the standard library for parsing CLI flags
	"flag"
	"net/http"
	"time"

	"github.com/solarwinds/golessons/web"
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
		fmt.Println("Happy to see you, Gophers!")
	} else {
		fmt.Println("Ugh, now I see Gophers.")
	}

	// it's easy to make a literal value - don't forget commas at the end
	// of *every* struct member line!
	metric := web.Metric{
		Name:      "cpu",
		Value:     float64(24.123),
		Timestamp: time.Now().UTC().UnixNano(),
	}

	fmt.Printf("Metric: %+v\n", metric)

	portString := fmt.Sprintf(":%d", port)
	muxAndServe(portString)
}

func muxAndServe(portString string) {
	fmt.Printf("binding to %s\n", portString)
	http.Handle("/hello", web.GetHello(friendly))
	http.ListenAndServe(portString, nil)
}
