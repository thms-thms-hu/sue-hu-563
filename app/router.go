package main

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/presentation/resource"
)

func initEndpoints() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST("/resource", resource.ApplyResource)

	router.Run(":12128")
}
