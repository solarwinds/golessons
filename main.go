package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"os/signal"

	"github.com/solarwinds/golessons/web"
)

// startTime will hold the time that this code executed and is used to calculate
// a runtime duration
var startTime = time.Now()

// sigIntChan will hold only a single value of type os.Signal and is used to catch
// SIGINT
var sigIntChan = make(chan os.Signal, 1)

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
	// signal package's Notify is used to tell the runtime what channel to put a
	// received signal value on
	signal.Notify(sigIntChan, os.Interrupt)
	go processSigInt()

	if friendly {
		fmt.Println("Happy to see you, Gophers!")
	} else {
		fmt.Println("Ugh, now I see Gophers.")
	}

	portString := fmt.Sprintf(":%d", port)
	fmt.Printf("binding to %s\n", portString)

	// here we register our GetHello function in the web package to handle GET /hello
	http.Handle("/hello", web.GetHello(friendly))

	// ListenAndServe with no arguments will use the http package's DefaultServeMux to
	// route requests to anything you've set up with http.Handle as above.
	http.ListenAndServe(portString, nil)
}

// processSigInt will block until the channel has a value
func processSigInt() {
	<-sigIntChan // NOTE: this blocks until there's something on the channel!

	fmt.Printf("\n[-] Caught interrupt\n")
	runDuration := time.Since(startTime)
	fmt.Printf("[*] Process ran for %v seconds\n", runDuration)
	os.Exit(0)
}
