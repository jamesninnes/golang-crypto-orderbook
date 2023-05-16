package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOrder(t *testing.T) {
	orderBook := NewOrderBook()

	order := &Order{
		ID:        1,
		AccountID: 1,
		OrderType: "buy",
		Price:     1000.0,
		Quantity:  10.0,
	}

	orderBook.AddOrder(order.AccountID, order.OrderType, order.Price, order.Quantity)

	assert.Contains(t, orderBook.BuyOrders, order.ID, "The order book should contain the added order")
}

func TestCancelOrder(t *testing.T) {
	orderBook := NewOrderBook()

	order := &Order{
		ID:        1,
		AccountID: 1,
		OrderType: "buy",
		Price:     1000.0,
		Quantity:  10.0,
	}

	orderBook.AddOrder(order.AccountID, order.OrderType, order.Price, order.Quantity)
	orderBook.CancelOrder(order.AccountID, order.ID)

	assert.NotContains(t, orderBook.BuyOrders, order.ID, "The order book should not contain the canceled order")
}
