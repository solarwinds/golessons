package main

import (
	"fmt"
	"os/signal"
	// flag is a package from the standard library for parsing CLI flags
	"flag"
	"net/http"
	"time"

	"os"

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

// processSigInt will block until the channel has a value
func processSigInt() {
	<-sigIntChan // NOTE: this blocks until there's something on the channel!

	fmt.Printf("\n[-] Caught interrupt\n")
	runDuration := time.Since(startTime)
	fmt.Printf("[*] Process ran for %v seconds\n", runDuration)
	os.Exit(0)
}

func muxAndServe(portString string) {
	fmt.Printf("binding to %s\n", portString)
	http.Handle("/hello", web.GetHello(friendly))
	http.ListenAndServe(portString, nil)
}
