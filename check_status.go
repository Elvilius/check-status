package check_status

import (
	"time"

	"github.com/Elvilius/check-status/internal/fetcher"
	"github.com/Elvilius/check-status/internal/interfaces"
	"github.com/Elvilius/check-status/internal/models"
	"github.com/Elvilius/check-status/internal/monitor"
	"github.com/Elvilius/check-status/internal/storage"
	"github.com/Elvilius/check-status/pkg/config"
)

type CheckStatus struct {
	storage  interfaces.Storage
	fetchers []*fetcher.Fetcher
	monitor  *monitor.Monitor
}

func NewCheckStatus(cfg []config.ProviderConfig) *CheckStatus {
	store := storage.NewMemoryStorage()
	monitor := monitor.NewMonitor()
	cs := &CheckStatus{
		storage: store,
		monitor: monitor,
	}

	for _, providerCfg := range cfg {
		f := fetcher.NewFetcher(providerCfg, store, monitor)
		f.Start()
		cs.fetchers = append(cs.fetchers, f)
	}
	return cs
}

func (cs *CheckStatus) GetOrderStatus(orderID int) (models.OrderStatus, error) {
	return cs.storage.Get(orderID)
}

func (cs *CheckStatus) GetMetrics() (int, int, time.Duration) {
	return cs.monitor.GetMetrics()
}

func (cs *CheckStatus) GetMessageMetric() string {
	return cs.monitor.GetMessageMetric()
}
