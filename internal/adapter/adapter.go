package adapter

import (
	"encoding/json"

	"github.com/Elvilius/check-status/internal/models"
)

type ProviderAdapter interface {
	AdaptResponse(data []byte) ([]models.OrderStatus, error)
}

type DefaultProviderAdapter struct{}

func (a *DefaultProviderAdapter) AdaptResponse(data []byte) ([]models.OrderStatus, error) {
	var response []struct {
		OrderID int    `json:"order_id"`
		Status  string `json:"status"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	length := len(response)
	orders := make([]models.OrderStatus, length)

	for i, order := range response {
		orders[i] = models.OrderStatus{
			OrderID: order.OrderID,
			Status:  order.Status,
		}
	}
	return orders, nil
}
