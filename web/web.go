package web

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// A struct lets us define a complex data type
type Metric struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

// PostMetric processes metrics
func PostMetric() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var m = &Metric{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(m); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("processing metric - %+v\n", m)

		// Simulate expensive operation with a half-second sleep
		time.Sleep(500 * time.Millisecond)

		w.WriteHeader(http.StatusAccepted)
	})
}

// GetHello's HTTP handler now creates and serializes a Metric
func GetHello(friendly bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric := NewMetric("cpu", 23.1234)

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

// NewMetric is a constructor for a Metric
func NewMetric(name string, value float64) *Metric {
	return &Metric{
		Name:      name,
		Value:     value,
		Timestamp: time.Now().UTC().UnixNano(),
	}
}

// MicrosecondTimestamp converts the timestamp to millisecond precision
func (m *Metric) MicrosecondTimestamp() int64 {
	return m.Timestamp / 1000
}

// ReduceTimestampPrecision changes the precision to be millisecond
func (m *Metric) ReduceTimestampPrecision() {
	m.Timestamp = m.MicrosecondTimestamp()
}
