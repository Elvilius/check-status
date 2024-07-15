package fetcher

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Elvilius/check-status/internal/config"
	"github.com/Elvilius/check-status/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	data map[int]models.OrderStatus
}

func (m *MockStorage) Save(orderID int, status models.OrderStatus) error {
	if m.data == nil {
		m.data = make(map[int]models.OrderStatus)
	}
	m.data[orderID] = status
	return nil
}

func (m *MockStorage) Get(orderID int) (models.OrderStatus, error) {
	status, ok := m.data[orderID]
	if !ok {
		return models.OrderStatus{}, errors.New("order not found")
	}
	return status, nil
}

func TestFetcher_fetchStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`[{"order_id": 12345, "status": "delivered"}]`))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))
	defer server.Close()

	fmt.Println(server)
	providerCfg := config.ProviderConfig{
		URL:         server.URL,
		Method:      http.MethodGet,
		Interval:    1 * time.Second,
		AuthHeaders: map[string]string{"Authorization": "Bearer token"},
	}

	mockStorage := &MockStorage{}

	f := NewFetcher(providerCfg, mockStorage)

	f.fetchStatus()

	status, err := mockStorage.Get(12345)
	assert.NoError(t, err)
	assert.Equal(t, "delivered", status.Status)
}
