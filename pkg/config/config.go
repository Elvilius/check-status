package config

import (
	"github.com/Elvilius/check-status/internal/adapter"
)

type ProviderConfig struct {
	URL         string
	Interval    int
	AuthHeaders map[string]string
	Method      string
	Adapter     adapter.ProviderAdapter
}
