package api

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Contact : Model for Contact
type ContactUser struct {
	Idcoopermapp string `json:"idcoopermapp" gorm:"primary_key"`
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

// CreateContactInput : struct for create contact post request
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

// FindContacts : Controller for getting all contacts
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []ContactUser

	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateContact : controller for creating new contact
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateContactUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your desired region
	}))
	svc := dynamodb.New(sess)
	tableName := "contato"
	id := uuid.New().String()
	putItemInput := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id":          {S: aws.String(id)},
			"name":        {S: aws.String(input.Name)},
			"email":       {S: aws.String(input.Email)},
			"cpf":         {S: aws.String(input.Cpf)},
			"cellphone":   {S: aws.String(input.Cellphone)},
			"cep":         {S: aws.String(input.Cep)},
			"cidade":      {S: aws.String(input.Cidade)},
			"uf":          {S: aws.String(input.Uf)},
			"endereco":    {S: aws.String(input.Endereco)},
			"complemento": {S: aws.String(input.Complemento)},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.PutItem(putItemInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create contact sqlite
	user := ContactUser{Idcoopermapp: id, Name: input.Name, Email: input.Email, Cpf: input.Cpf, Cellphone: input.Cellphone, Cep: input.Cep, Cidade: input.Cidade, Uf: input.Uf, Endereco: input.Endereco, Complemento: input.Complemento}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func HealthCheck(c *gin.Context) {
	// Check the health of the server and return a status code accordingly
	if ServerIsHealthy() {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "not ok"})
	}

}

func ServerIsHealthy() bool {
	// Check the health of the server and return true or false accordingly
	// For example, check if the server can connect to the database
	return true
}
