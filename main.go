package main

import (
	"contactform/api"
	"contactform/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// create/configure database instance
	db := config.CreateDatabase()
	// set db to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// routes definition for finding and creating contacts
	r.GET("/health", api.HealthCheck)
	r.GET("/contacts", api.FindUsers)
	r.POST("/contact", api.CreateUser)
	// start the server
	r.SetTrustedProxies(nil)

	r.Run(":8080")

}
