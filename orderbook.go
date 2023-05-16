package main

import (
	"sync"
	"time"
)

type Order struct {
	ID         int64
	AccountID  int64
	OrderType  string
	Price      float64
	Quantity   float64
	Timestamp  time.Time
}

type OrderBook struct {
	BuyOrders  map[int64]*Order
	SellOrders map[int64]*Order
	sync.RWMutex
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		BuyOrders:  make(map[int64]*Order),
		SellOrders: make(map[int64]*Order),
	}
}

func (o *OrderBook) AddOrder(order *Order) {
	if order.OrderType == "buy" {
		o.BuyOrders[order.ID] = order
	} else if order.OrderType == "sell" {
		o.SellOrders[order.ID] = order
	}
}

func (o *OrderBook) CancelOrder(orderID, accountID int64, orderType string) bool {
	if orderType == "buy" {
		order, exists := o.BuyOrders[orderID]
		if !exists || order.AccountID != accountID {
			return false
		}
		delete(o.BuyOrders, orderID)
	} else if orderType == "sell" {
		order, exists := o.SellOrders[orderID]
		if !exists || order.AccountID != accountID {
			return false
		}
		delete(o.SellOrders, orderID)
	}
	return true
}

func (o *OrderBook) ModifyOrder(orderID, accountID int64, newPrice, newQuantity float64) bool {
	order, exists := o.BuyOrders[orderID]
	if exists && order.AccountID == accountID {
		order.Price = newPrice
		order.Quantity = newQuantity
		return true
	}

	order, exists = o.SellOrders[orderID]
	if exists && order.AccountID == accountID {
		order.Price = newPrice
		order.Quantity = newQuantity
		return true
	}

	return false
}

func (o *OrderBook) GetOrderType(orderID int64) string {
	o.RLock()
	defer o.RUnlock()

	if _, exists := o.BuyOrders[orderID]; exists {
		return "buy"
	}
	if _, exists := o.SellOrders[orderID]; exists {
		return "sell"
	}
	return ""
}

var orderBook = NewOrderBook()
