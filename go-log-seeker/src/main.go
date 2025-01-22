package main

import (
	"github.com/gin-gonic/gin"
	"go-log-seeker/src/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
