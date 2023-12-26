package main

import (
	"contactform/api"
	"contactform/config"
	"contactform/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080", "http://coopermapp.aplicativos.rio.br", "https://coopermapp.aplicativos.rio.br", "https://dynamodb.us-east-1.amazonaws.com", "https://coopermappoffice.aplicativos.rio.br"},
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
	r.GET("/viacep/:cep", utils.FindAdressByZipCode)
	r.GET("/fipe/marcas/:tipo/:codigo", utils.GetMarca)
	r.GET("/fipe/modelos/:tipo/:marca/:codigo", utils.GetModelo)
	r.GET("/contacts", api.FindContacts)
	r.GET("/getendereco/:cpf/:idcoopermapp", api.FindContactsEnderecoByCpfAndId)
	r.GET("/contactsendereco", api.FindContactsEndereco)
	r.POST("/contact", api.CreateUser)
	r.POST("/contactendereco", api.CreateUserEndereco)
	r.POST("/contactveiculo", api.CreateUserVeiculo)
	r.GET("/getveiculo/:cpf/:idcoopermapp", api.FindContactsVeiculoByCpfAndId)
	r.GET("/contactsveiculo", api.FindContactsVeiculo)

	// start the server
	r.SetTrustedProxies(nil)
	r.Run(":8080")

}
