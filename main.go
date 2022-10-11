package main

import (
	"PresentationManagement-backend/controllers"
	"PresentationManagement-backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Presentation Management Systems API Online."})
	})

	public := r.Group("/api")

	public.POST("/presentation/", controllers.CreatePresentation)
	public.GET("/presentation/", controllers.GetPresentation)
	public.PUT("/presentation/:id", controllers.UpdatePresentation)
	public.DELETE("/presentation/:id", controllers.DeletePresentation)

	public.PUT("/presentation/pp/:id", controllers.UpdatePowerpoint)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server.")
		return
	}
}
