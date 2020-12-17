// main.go

package main

import (
	"authapp/controllers"
	"authapp/database"
	"authapp/middlewares"
	"authapp/models"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		// here
		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			protected.GET("/profile", controllers.Profile)
		}
	}

	return r
}

func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.GlobalDB.AutoMigrate(&models.User{})

	r := setupRouter()
	r.Run(":8080")
}
