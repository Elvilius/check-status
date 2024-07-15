package storage

import (
	"testing"

	"github.com/Elvilius/check-status/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStorage_Save(t *testing.T) {
	tests := []struct {
		name    string
		initial map[int]models.OrderStatus
		orderID int
		status  models.OrderStatus
		wantErr bool
	}{
		{
			name:    "Save new order status",
			initial: map[int]models.OrderStatus{},
			orderID: 1,
			status: models.OrderStatus{
				OrderID: 1,
				Status:  "processing",
			},
			wantErr: false,
		},
		{
			name: "Update existing order status",
			initial: map[int]models.OrderStatus{
				1: {OrderID: 1, Status: "pending"},
			},
			orderID: 1,
			status: models.OrderStatus{
				OrderID: 1,
				Status:  "shipped",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMemoryStorage()
			ms.data = tt.initial

			err := ms.Save(tt.orderID, tt.status)

			assert.Equal(t, tt.wantErr, err != nil)

			got, ok := ms.data[tt.orderID]
			assert.True(t, ok)
			assert.Equal(t, tt.status, got)
		})
	}
}

func TestMemoryStorage_Get(t *testing.T) {
	tests := []struct {
		name    string
		initial map[int]models.OrderStatus
		orderID int
		want    models.OrderStatus
		wantErr bool
	}{
		{
			name: "Get existing order status",
			initial: map[int]models.OrderStatus{
				1: {OrderID: 1, Status: "processing"},
			},
			orderID: 1,
			want: models.OrderStatus{
				OrderID: 1,
				Status:  "processing",
			},
			wantErr: false,
		},
		{
			name:    "Get non-existing order status",
			initial: map[int]models.OrderStatus{},
			orderID: 2,
			want:    models.OrderStatus{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := NewMemoryStorage()
			ms.data = tt.initial

			got, err := ms.Get(tt.orderID)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
