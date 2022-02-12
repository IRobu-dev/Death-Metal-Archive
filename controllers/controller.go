package controllers

import (
	"archive-be/services"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	router := gin.Default()
	router.GET("/hello", services.GetHello)
	router.GET("/albums", services.GetAlbums)
	router.GET("/band/:key", services.SearchBandWithAlbums)
	router.GET("/album/:key", services.GetBandsFromAlbums)
	router.POST("/data", services.GetDataFromCSV)
	router.Run("localhost:8080")
}
