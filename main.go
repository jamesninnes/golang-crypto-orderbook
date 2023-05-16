package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterOrderBookRoutes(r)
	RegisterAccountRoutes(r)

	r.Run()
}
