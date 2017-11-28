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
func GetHello(verbose bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		metric := Metric{
			Name:      "cpu",
			Value:     24.1234,
			Timestamp: time.Now().UTC().UnixNano(),
		}

		if verbose {
			bodyData, err := json.Marshal(&metric)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(bodyData)
			return
		} else {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("We've got a metric"))
		}
	})
}
