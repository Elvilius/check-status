# check-status

`check-status` is a Go library designed to periodically fetch and store order statuses from various providers. This library helps in managing the status of orders in memory and provides an interface to retrieve these statuses.

## Features

- Fetch order statuses from multiple providers at specified intervals.
- Store order statuses in memory.
- Retrieve order statuses by order ID.

## Installation

To install the library, run:

```sh
go get github.com/Elvilius/check-status
```

## Usage
Importing the library


```go
import (
	"github.com/Elvilius/check-status"
)
```

## Initializing CheckStatus

To initialize CheckStatus, you need to provide a slice of config.ProviderConfig which contains the configuration for each provider

```go
package main

import (
	"github.com/Elvilius/check-status"
	"github.com/Elvilius/check-status/pkg/config"
)

func main() {
	providerConfigs := []config.ProviderConfig{
		{
			URL: "http://example.com/api/order-status",
			Method: "GET",
			Interval: 10,
			AuthHeaders: map[string]string{
				"Authorization": "Bearer your_token",
			},
			Adapter: &YourProviderAdapter{},
		},
		// Add more provider configurations as needed
	}

	cs := check_status.NewCheckStatus(providerConfigs)

	time.Sleep(10 * time.Second)
	// Use cs to get order status
	orderID := 12345
	status, err := cs.GetOrderStatus(orderID)
	if err != nil {
		// Handle error
	}
	// Use status
}
```

By default, the library uses the following adapter:
```go
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
```

### Configuring Providers
Each ProviderConfig contains the following fields:

- URL: The URL of the provider's API endpoint.
- Method: The HTTP method to use (e.g., "GET", "POST").
- Interval: The interval at which the fetcher should request the provider's API.
- AuthHeaders: A map of authentication headers required by the provider's API.
- Adapter: An implementation of the Adapter interface to adapt the provider's response to models.OrderStatus.
- Implementing an Adapter
You need to implement an adapter that converts the provider's response into the models.OrderStatus format. Here's an example



## Implementing an Adapter
You need to implement an adapter that converts the provider's response into the models.OrderStatus format. Here's an example:

```go
type YourProviderAdapter struct{}

func (a *YourProviderAdapter) AdaptResponse(data []byte) ([]models.OrderStatus, error) {
	// Implement the logic to adapt the provider's response
}
```
## Fetching Order Status
Once CheckStatus is initialized and running, you can fetch the status of an order using its ID
