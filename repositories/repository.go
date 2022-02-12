package repositories

import (
	"archive-be/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// CARICO PRIMA TUTTI I CSV IN UN QUARTO METODO. CARICO TUTTI I FILE ALL INIZIO DELL'APP IN MODO CHE NON FACCIA RITARDI.

func LoadBands() []models.Band {
	csvFile, err := os.Open("src/bands.csv")
	if err != nil {
		fmt.Printf("Bands file exist? %v \n", os.IsExist(err))
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var bandsList []models.Band
	for _, line := range csvLines {
		idS, _ := strconv.Atoi(line[0])
		formedInS, _ := strconv.Atoi(line[4])
		band := models.Band{
			Id:       idS,
			Name:     line[1],
			Country:  line[2],
			Status:   line[3],
			FormedIn: formedInS,
			Genre:    line[5],
			Theme:    line[6],
			Active:   line[7],
		}
		bandsList = append(bandsList, band)
	}
	bandsList = bandsList[1:]
	return bandsList
}

func LoadAlbums() []models.Album {
	csvFile, err := os.Open("src/albums.csv")
	if err != nil {
		fmt.Printf("Album file exist? %v \n", os.IsExist(err))
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var albumList []models.Album
	for _, line := range csvLines {
		idS, _ := strconv.Atoi(line[0])
		bandS, _ := strconv.Atoi(line[1])
		yearS, _ := strconv.Atoi(line[3])
		album := models.Album{
			Id:    idS,
			Band:  bandS,
			Title: line[2],
			Year:  yearS,
		}
		albumList = append(albumList, album)
	}
	albumList = albumList[1:]
	return albumList
}

func LoadReview() []models.Reviews {
	csvFile, err := os.Open("src/reviews.csv")
	if err != nil {
		fmt.Printf("Review file exist? %v \n", os.IsExist(err))
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var reviewList []models.Reviews
	for _, line := range csvLines {
		idS, _ := strconv.Atoi(line[0])
		scoreS, _ := strconv.ParseFloat(line[3], 32)
		review := models.Reviews{
			Id:      idS,
			Album:   line[1],
			Title:   line[2],
			Score:   float32(scoreS),
			Content: line[4],
		}
		reviewList = append(reviewList, review)
	}
	reviewList = reviewList[1:]
	return reviewList
}
