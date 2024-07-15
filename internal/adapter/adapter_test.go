package adapter

import (
	"encoding/xml"
	"errors"
	"testing"

	"github.com/Elvilius/check-status/internal/models"
	"github.com/stretchr/testify/assert"
)

type XMLProviderAdapter struct{}

func (c *XMLProviderAdapter) AdaptResponse(data []byte) ([]models.OrderStatus, error) {
	var orderXML struct {
		OrderID int    `xml:"OrderID"`
		Status  string `xml:"Status"`
	}

	if err := xml.Unmarshal(data, &orderXML); err != nil {
		return nil, err
	}

	if orderXML.OrderID == 0 || orderXML.Status == "" {
		return nil, errors.New("missing required fields in XML")
	}

	return []models.OrderStatus{
		{
			OrderID: orderXML.OrderID,
			Status:  orderXML.Status,
		},
	}, nil
}

func TestDefaultProviderAdapter_AdaptResponse(t *testing.T) {
	tests := []struct {
		name    string
		a       *DefaultProviderAdapter
		data    []byte
		want    []models.OrderStatus
		wantErr bool
	}{
		{
			name: "Valid JSON array data",
			a:    &DefaultProviderAdapter{},
			data: []byte(`[{"order_id": 12345, "status": "delivered"}]`),
			want: []models.OrderStatus{
				{
					OrderID: 12345,
					Status:  "delivered",
				},
			},
			wantErr: false,
		},
		{
			name: "Valid JSON single object data",
			a:    &DefaultProviderAdapter{},
			data: []byte(`[{"order_id": 54321, "status": "pending"}]`),
			want: []models.OrderStatus{
				{
					OrderID: 54321,
					Status:  "pending",
				},
			},
			wantErr: false,
		},
		{
			name:    "Invalid JSON data",
			a:       &DefaultProviderAdapter{},
			data:    []byte(`invalid json`),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Empty data",
			a:       &DefaultProviderAdapter{},
			data:    []byte(``),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.AdaptResponse(tt.data)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestXMLProviderAdapter_AdaptResponse(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		a       *XMLProviderAdapter
		args    args
		want    []models.OrderStatus
		wantErr bool
	}{
		{
			name: "Valid XML data",
			a:    &XMLProviderAdapter{},
			args: args{
				data: []byte(`<Order><OrderID>12345</OrderID><Status>delivered</Status></Order>`),
			},
			want: []models.OrderStatus{
				{
					OrderID: 12345,
					Status:  "delivered",
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid XML data",
			a:    &XMLProviderAdapter{},
			args: args{
				data: []byte(`invalid xml`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Empty data",
			a:    &XMLProviderAdapter{},
			args: args{
				data: []byte(``),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.AdaptResponse(tt.args.data)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
