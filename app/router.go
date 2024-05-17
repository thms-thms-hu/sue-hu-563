package main

import (
	"github.com/gin-gonic/gin"
	"sue-hu-563/presentation"
)

func initEndpoints() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.POST("/resource", presentation.ApplyResource)
}
