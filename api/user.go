package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Artist : Model for artist
type ContactUser struct {
	Idcoopermapp uint   `json:"idcoopermapp" gorm:"primary_key"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique;not null"`
	Cpf          string `json:"cpf" gorm:"unique;not null"`
	Cellphone    string `json:"cellphone" gorm:"unique;not null"`
	Cep          string `json:"cep"`
	Cidade       string `json:"cidade"`
	Uf           string `json:"uf"`
	Endereco     string `json:"endereco"`
	Complemento  string `json:"complemento"`
}

// CreateArtistInput : struct for create art post request
type CreateContactUserInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Cpf         string `json:"cpf" binding:"required"`
	Cellphone   string `json:"cellphone" binding:"required"`
	Cep         string `json:"cep"`
	Cidade      string `json:"cidade"`
	Uf          string `json:"uf"`
	Endereco    string `json:"endereco"`
	Complemento string `json:"complemento"`
}

// FindArtists : Controller for getting all artists
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []ContactUser

	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateArtist : controller for creating new artists
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateContactUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create artist
	user := ContactUser{Name: input.Name, Email: input.Email}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func healthCheck(c *gin.Context) {
	// Check the health of the server and return a status code accordingly
	if serverIsHealthy() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
	c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unavailable"})
}

func serverIsHealthy() bool {
	// Check the health of the server and return true or false accordingly
	// For example, check if the server can connect to the database
	return true
}
