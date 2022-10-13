package controllers

import (
	"PresentationManagement-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type SettingInput struct {
	Company   string `json:"company" binding:"required"`
	EventName string `json:"event" binding:"required"`
}

func CreateSettings(c *gin.Context) {
	var input SettingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Settings{}
	s.Company = input.Company
	s.EventName = input.EventName

	_, err := s.CreateSettings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings created."})
}

func GetSettings(c *gin.Context) {
	settings, err := models.GetSettings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"settings": settings})
}

func UpdateSettings(c *gin.Context) {
	var input SettingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Settings{}
	s.ID = uint(id)
	s.Company = input.Company
	s.EventName = input.EventName

	_, err = s.UpdateSettings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated."})
}

func DeleteSettings(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := models.Settings{}
	s.ID = uint(id)

	_, err = s.DeleteSettings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Settings deleted."})
}

func UploadLogo(c *gin.Context) {

	//BLOCK - File upload & rename
	file, err := c.FormFile("file")
	if err != nil {
		//File has not been recieved
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileExtension := filepath.Ext(file.Filename)
	newFileName := "logo" + fileExtension
	if err := c.SaveUploadedFile(file, os.Getenv("BUCKET_PATH")+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//END BLOCK

	c.JSON(http.StatusOK, gin.H{"message": "Logo uploaded."})
}

func GetSponsors(c *gin.Context) {
	output := []string{}

	files, err := ioutil.ReadDir("./storage/sponsors/")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for _, file := range files {
		if !file.IsDir() {
			output = append(output, file.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": output})
}

func UploadSponsor(c *gin.Context) {
	//BLOCK - File upload & rename
	file, err := c.FormFile("file")
	if err != nil {
		//File has not been recieved
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileExtension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + fileExtension
	if err := c.SaveUploadedFile(file, os.Getenv("BUCKET_PATH")+"sponsors/"+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//END BLOCK

	c.JSON(http.StatusOK, gin.H{"message": "Sponsor Logo uploaded."})
}
