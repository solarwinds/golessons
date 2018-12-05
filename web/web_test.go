package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/solarwinds/golessons/processors"
)

// TestMetric_MillisecondTimestamp is run automatically because it starts with "Test"
func TestMetric_MillisecondTimestamp(t *testing.T) {
	in := int64(1511898670850117031)
	out := int64(1511898670850117)

	metric := &processors.Metric{
		Name:      "test",
		Value:     2.71828,
		Timestamp: in,
	}

	result := metric.MicrosecondTimestamp()

	if result != out {
		t.Errorf("expected %d to equal %d but it didn't", result, out)
	}
}

func TestPostMetric_MockProcessor(t *testing.T) {
	testChan := make(chan *processors.Metric)
	testProcessor := processors.NewMockMetricProcessor()
	testMetric := &processors.Metric{
		Name:      "Test",
		Value:     1.23,
		Timestamp: 1544046763,
	}
	jsonMetric, _ := json.Marshal(testMetric)

	req, err := http.NewRequest("POST", "/metrics", bytes.NewBuffer(jsonMetric))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := PostMetric(testChan, testProcessor)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}
}
