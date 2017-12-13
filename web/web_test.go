package web

import "testing"

// TestMetric_MillisecondTimestamp is run automatically because it starts with "Test"
func TestMetric_MillisecondTimestamp(t *testing.T) {
	in := int64(1511898670850117031)
	out := int64(1511898670850117)

	metric := &Metric{
		Name:      "test",
		Value:     2.71828,
		Timestamp: in,
	}

	result := metric.MicrosecondTimestamp()

	if result != out {
		t.Errorf("expected %d to equal %d but it didn't", result, out)
	}
}
