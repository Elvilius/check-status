package fetcher

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Elvilius/check-status/internal/adapter"
	"github.com/Elvilius/check-status/internal/config"
	"github.com/Elvilius/check-status/internal/interfaces"
)

type Fetcher struct {
	providerCfg config.ProviderConfig
	storage     interfaces.Storage
	ticker      *time.Ticker
}

func NewFetcher(providerCfg config.ProviderConfig, storage interfaces.Storage) *Fetcher {
	if providerCfg.Adapter == nil {
		providerCfg.Adapter = &adapter.DefaultProviderAdapter{}
	}

	return &Fetcher{
		providerCfg: providerCfg,
		storage:     storage,
		ticker:      time.NewTicker(time.Duration(providerCfg.Interval) * time.Second),
	}
}

func (f *Fetcher) Start() {
	go func() {
		for range f.ticker.C {
			f.fetchStatus()
		}
	}()
}

func (f *Fetcher) fetchStatus() {
	req, err := http.NewRequest(f.providerCfg.Method, f.providerCfg.URL, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	for key, value := range f.providerCfg.AuthHeaders {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response:", err)
		return
	}

	statuses, err := f.providerCfg.Adapter.AdaptResponse(body)
	if err != nil {
		log.Println("Error adapting response:", err)
		return
	}

	for _, status := range statuses {
		if err := f.storage.Save(status.OrderID, status); err != nil {
			log.Println("Error saving status:", err)
			continue
		}
	}
}
