package web

import (
	"encoding/json"
	"net/http"
	"time"
)

// A struct lets us define a complex data type
type Metric struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
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
