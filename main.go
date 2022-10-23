package main

import (
	"PresentationManagement-backend/controllers"
	"PresentationManagement-backend/models"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	models.ConnectDatabase()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{os.Getenv("APP_URL")}
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

	public.PUT("/presentation/pp/:session_id", controllers.UpdatePowerpoint)

	public.POST("/settings/", controllers.CreateSettings)
	public.GET("/settings/", controllers.GetSettings)
	public.PUT("/settings/:id", controllers.UpdateSettings)
	public.DELETE("/settings/:id", controllers.DeleteSettings)

	public.PUT("/settings/logo/", controllers.UploadLogo)
	public.PUT("/settings/sponsor/", controllers.UploadSponsor)
	public.GET("/settings/sponsor/", controllers.GetSponsors)


	//Serve Static Files
	public.StaticFS("/images", http.Dir("./storage"))

	err = r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		fmt.Println("Error starting server.")
		return
	}
}
