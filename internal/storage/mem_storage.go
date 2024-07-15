package storage

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Elvilius/check-status/internal/models"
)

type MemoryStorage struct {
	data map[int]models.OrderStatus
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]models.OrderStatus),
	}
}

func (ms *MemoryStorage) Save(orderID int, status models.OrderStatus) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.data[orderID] = status

	fmt.Println(ms.data)
	return nil
}

func (ms *MemoryStorage) Get(orderID int) (models.OrderStatus, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	status, ok := ms.data[orderID]
	if !ok {
		return models.OrderStatus{}, errors.New("order not found")
	}

	return status, nil
}
