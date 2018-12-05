package processors

import (
	"fmt"
)

type MockMetricProcessor struct{}

func NewMockMetricProcessor() *MockMetricProcessor {
	return &MockMetricProcessor{}
}

func (mmp *MockMetricProcessor) ProcessMetric(m *Metric) error {
	fmt.Println("Mocking the processing of a metric")
	return nil
}
