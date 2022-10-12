package main

import (
	"PresentationManagement-backend/controllers"
	"PresentationManagement-backend/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	models.ConnectDatabase()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080"}
	corsConfig.AddAllowMethods("OPTIONS")

	r := gin.Default()

	r.Use(cors.New(corsConfig))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Presentation Management Systems API Online."})
	})

	public := r.Group("/api")

	public.POST("/presentation/", controllers.CreatePresentation)
	public.GET("/presentation/", controllers.GetPresentation)
	public.PUT("/presentation/:id", controllers.UpdatePresentation)
	public.DELETE("/presentation/:id", controllers.DeletePresentation)

	public.PUT("/presentation/pp/:id", controllers.UpdatePowerpoint)

	err := r.Run(":8085")
	if err != nil {
		fmt.Println("Error starting server.")
		return
	}
}
