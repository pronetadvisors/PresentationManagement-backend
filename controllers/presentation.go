package controllers

import (
	"PresentationManagement-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type PresentationInput struct {
	SessionID  string    `json:"session_id" binding:"required"`
	Time       time.Time `json:"time" binding:"required"`
	EndTime    time.Time `json:"endtime" binding:"required"`
	Location   string    `json:"location" binding:"required"`
	Speaker    string    `json:"speaker" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Powerpoint string    `json:"powerpoint"`
}

func CreatePresentation(c *gin.Context) {
	var input PresentationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Presentation{}
	p.SessionID = input.SessionID
	p.Time = input.Time
	p.EndTime = input.EndTime
	p.Location = input.Location
	p.Speaker = input.Speaker
	p.Title = input.Title

	_, err := p.CreatePresentation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Presentation created."})
}

func GetPresentation(c *gin.Context) {
	presentations, err := models.GetPresentation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"presentations": presentations})
}

func UpdatePresentation(c *gin.Context) {
	var input PresentationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Presentation{}
	p.ID = uint(id)
	p.SessionID = input.SessionID
	p.Time = input.Time
	p.EndTime = input.EndTime
	p.Location = input.Location
	p.Speaker = input.Speaker
	p.Title = input.Title

	_, err = p.UpdatePresentation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Presentation updated."})
}

func DeletePresentation(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Presentation{}
	p.ID = uint(id)

	_, err = p.DeletePresentation()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Presentation deleted."})
}

func UpdatePowerpoint(c *gin.Context) {
	sessionId := c.Param("session_id")

	//BLOCK - File upload & rename
	file, err := c.FormFile("file")
	if err != nil {
		//File has not been recieved
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fileExtension := filepath.Ext(file.Filename)
	newFileName := sessionId + fileExtension
	if err := c.SaveUploadedFile(file, os.Getenv("BUCKET_PATH")+newFileName); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//END BLOCK

	p := models.Presentation{}
	p.SessionID = sessionId
	p.Powerpoint = newFileName

	_, err = p.UpdatePowerpoint()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Presentation uploaded and updated."})
}
