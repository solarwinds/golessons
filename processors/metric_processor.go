package processors

import (
	"errors"
	"time"
)

type Processor interface {
	ProcessMetric(*Metric) error
}

// A struct lets us define a complex data type
type Metric struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
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

// AsyncMetricProcessor processes metrics asynchronously
type AsyncMetricProcessor struct {
	metricsChan chan *Metric
}

func NewAsyncMetricProcessor(metricsChan chan *Metric) *AsyncMetricProcessor {
	return &AsyncMetricProcessor{
		metricsChan: metricsChan,
	}
}

// ProcessMetric sends the metric off on its channel if it has one,
// returns an error if channel doesn't exist
func (amp *AsyncMetricProcessor) ProcessMetric(m *Metric) error {
	if amp.metricsChan == nil {
		return errors.New("AsyncMetricProcessor requires non-nil metricsChan")
	}
	amp.metricsChan <- m
	return nil
}
