package monitor

import (
	"fmt"
	"sync"
	"time"
)

type Monitor struct {
	mu           sync.Mutex
	successCount int
	errorCount   int
	totalTime    time.Duration
}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) LogSuccess(duration time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.successCount++
	m.totalTime += duration
}

func (m *Monitor) LogError() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.errorCount++
}

func (m *Monitor) GetMetrics() (int, int, time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.successCount, m.errorCount, m.totalTime
}

func (m *Monitor) GetMessageMetric() string {
	success, errors, totalTime := m.GetMetrics()
	return fmt.Sprintf("Success: %d, Errors: %d, Total Time: %s", success, errors, totalTime)
}
