package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	RegisterOrderBookRoutes(r)
	RegisterAccountRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
