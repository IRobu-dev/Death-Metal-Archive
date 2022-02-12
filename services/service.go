package services

import (
	"archive-be/models"
	"archive-be/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

var albumsList = repositories.LoadAlbums()
var bandsList = repositories.LoadBands()
var reviews = repositories.LoadReview()

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albumsList)
}
func GetHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello World")
}

func GetDataFromCSV(context *gin.Context) {
	body := models.Pagination{}

	// using BindJson method to serialize body with struct
	//https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if body.CSV == "albumsList" {
		result := albumsList[body.First:body.Last]
		context.JSON(http.StatusAccepted, result)
	} else if body.CSV == "bandsList" {
		result := bandsList[body.First:body.Last]
		context.JSON(http.StatusAccepted, result)
	} else if body.CSV == "reviews" {
		result := reviews[body.First:body.Last]
		context.JSON(http.StatusAccepted, result)
	}
}

func SearchBandWithAlbums(context *gin.Context) {
	keyWord := context.Param("key")
	var ResultSearch models.SearchResult

	for _, band := range bandsList {
		if band.Name == keyWord {
			ResultSearch.Bands = append(ResultSearch.Bands, band)
			ResultSearch.Albums = append(ResultSearch.Albums, GetAlbumFromBands(band.Id)...)
		}
	}
	context.JSON(http.StatusAccepted, ResultSearch)
}

func GetAlbumFromBands(bandId int) []models.Album {
	var result []models.Album
	for _, album := range albumsList {
		if album.Band == bandId {
			result = append(result, album)
		}
	}
	return result
}
func GetBandsFromAlbums() {}
