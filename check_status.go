package check_status

import (
	"github.com/Elvilius/check-status/internal/config"
	"github.com/Elvilius/check-status/internal/fetcher"
	"github.com/Elvilius/check-status/internal/interfaces"
	"github.com/Elvilius/check-status/internal/models"
	"github.com/Elvilius/check-status/internal/storage"
)

type CheckStatus struct {
	storage  interfaces.Storage
	fetchers []*fetcher.Fetcher
}

func NewCheckStatus(cfg []config.ProviderConfig) *CheckStatus {
	store := storage.NewMemoryStorage()

	cs := &CheckStatus{
		storage: store,
	}

	for _, providerCfg := range cfg {
		f := fetcher.NewFetcher(providerCfg, store)
		f.Start()
		cs.fetchers = append(cs.fetchers, f)
	}
	return cs
}

func (cs *CheckStatus) GetOrderStatus(orderID int) (models.OrderStatus, error) {
	return cs.storage.Get(orderID)
}
