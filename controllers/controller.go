package controllers

import (
	"archive-be/services"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()
	router.GET("/hello", services.GetHello)
	router.GET("/albums", services.GetAlbums)
	router.GET("/data/:key", services.SearchBandWithAlbums)
	router.POST("/data", services.GetDataFromCSV)
	router.Run("localhost:8080")
}
