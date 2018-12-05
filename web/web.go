package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/solarwinds/golessons/processors"
)

var stopProcessing bool

// PostMetric processes metrics
func PostMetric(metricsChan chan<- *processors.Metric, metricProcessor processors.Processor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !stopProcessing {
			var m = &processors.Metric{}
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(m); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err := metricProcessor.ProcessMetric(m)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
			}
			w.WriteHeader(http.StatusAccepted)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

// ProcessMetrics takes a Metrics channel and a control channel and implements processing logic
// until told to stop
func ProcessMetrics(metricsChan <-chan *processors.Metric, stopChan <-chan bool) {
P:
	for {
		select {
		case m := <-metricsChan:
			log.Printf("processing metric - %+v\n", m)
		case <-stopChan:
			stopProcessing = true
			break P
		}
	}
}

// GetHello's HTTP handler now creates and serializes a Metric
func GetHello(friendly bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric := processors.NewMetric("cpu", 23.1234)

		// reduce timestamp precision if `friendly` is true
		if friendly {
			metric.ReduceTimestampPrecision()
		}
		bodyData, err := json.Marshal(&metric)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bodyData)
	})
}
