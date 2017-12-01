package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/signal"

	"os"
	"time"

	"github.com/solarwinds/golessons/web"
)

// startTime will hold the time that this code executed and is used to calculate
// a runtime duration
var startTime = time.Now()

var sigIntChan = make(chan os.Signal, 1)
var metricsChan = make(chan *web.Metric)
var stopChan = make(chan bool)

// friendly is a variable that is in scope everywhere in the main package
var friendly bool

// port will hold the value of the TCP port we wish to bind to
var port int

// timeLimit will hold the value of how many metrics to accept
var timeLimit int

// init is a special function like main and is guaranteed to run before main
func init() {
	// This sets up a boolean flag, held in the `friendly` variable so that it's accessible everywhere
	flag.BoolVar(&friendly, "friendly", false, "whether to print out a greeting")

	flag.IntVar(&port, "port", 7777, "the TCP port to bind to")
	flag.IntVar(&timeLimit, "time-limit", 0, "limit the number of seconds to run")

	flag.Parse()
}

func main() {
	signal.Notify(sigIntChan, os.Interrupt)
	go processSigInt()

	if friendly {
		fmt.Println("Happy to see you, Gophers!")
	} else {
		fmt.Println("Ugh, now I see Gophers.")
	}

	// handle processing of metrics
	go web.ProcessMetrics(metricsChan, stopChan)

	if timeLimit > 0 {
		go turnOffWhenTimeExpires()
	}

	portString := fmt.Sprintf(":%d", port)
	muxAndServe(portString)
}

func turnOffWhenTimeExpires() {
	for {
		if time.Since(startTime) > time.Duration(timeLimit)*time.Second {
			stopChan <- true
			break
		}
	}
	close(stopChan)
	close(metricsChan)
}

// processSigInt will block until the channel has a value
func processSigInt() {
	<-sigIntChan // NOTE: this blocks until there's something on the channel!

	close(metricsChan)
	close(stopChan)

	fmt.Printf("\n[-] Caught interrupt\n")
	runDuration := time.Since(startTime)
	fmt.Printf("[*] Process ran for %v seconds\n", runDuration)
	os.Exit(0)
}

func muxAndServe(portString string) {
	fmt.Printf("binding to %s\n", portString)
	http.Handle("/hello", web.GetHello(friendly))
	http.Handle("/metrics", web.PostMetric(metricsChan))
	http.ListenAndServe(portString, nil)
}
