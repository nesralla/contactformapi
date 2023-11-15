package main

import (
	"contactform/api"
	"contactform/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://coopermapp.aplicativos.rio.br", "https://coopermapp.aplicativos.rio.br"},
		AllowMethods: []string{"POST", "HEAD", "PATCH", "GET", "PUT"},
		AllowHeaders: []string{"Origin"},
		MaxAge:       12 * time.Hour,
	}))
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
	r.POST("/contactendereco", api.CreateUserEndereco)
	r.POST("/contactveiculo", api.CreateUserVeiculo)
	// start the server
	r.SetTrustedProxies(nil)

	autotls.Run(r, "contact.api.rio.br", "*.api.rio.br", "coopermapp.aplicativos.rio.br", "*.aplicativos.rio.br")

}
