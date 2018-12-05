package web

import (
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

	req, err := http.NewRequest("POST", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := PostMetric(testChan, testProcessor)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
