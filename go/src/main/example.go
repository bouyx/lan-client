package main

import (
	"fmt"
	"main/ressources"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("nique")
	r := gin.Default()
	r.GET("/quiche", ressources.Quiche)
	r.GET("/couille", ressources.Couille)

	r.Run() // listen and serve on 0.0.0.0:8080
}
