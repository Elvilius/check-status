package interfaces

import "github.com/Elvilius/check-status/internal/models"

type Storage interface {
	Save(orderID int, status models.OrderStatus) error
	Get(orderID int) (models.OrderStatus, error)
}
