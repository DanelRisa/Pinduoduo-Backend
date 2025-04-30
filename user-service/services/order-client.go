package services

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Order struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}

func GetOrdersByUserID(userID uint) ([]Order, error) {
	client := resty.New()

	var orders []Order
	resp, err := client.R().
		SetQueryParam("user_id", fmt.Sprintf("%d", userID)).
		SetResult(&orders).
		Get("http://localhost:8082/orders")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("order-service error: %s", resp.Status())
	}

	return orders, nil
}
