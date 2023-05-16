package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterOrderBookRoutes(r *gin.Engine) {
	r.POST("/api/add_order", handleAddOrder)
	r.POST("/api/cancel_order", handleCancelOrder)
	r.POST("/api/modify_order", handleModifyOrder)
	r.GET("/api/orderbook", handleGetOrderBook)
}
