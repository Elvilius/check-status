package config

import (
	"time"

	"github.com/Elvilius/check-status/internal/adapter"
)

type ProviderConfig struct {
	URL         string
	Interval    time.Duration
	AuthHeaders map[string]string
	Method      string
	Adapter     adapter.ProviderAdapter
}
