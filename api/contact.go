package api

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Contact : Model for Contact
type ContactUser struct {
	ID           uint   `gorm:"primary_key;autoincrement"`
	Idcoopermapp string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Documento    string `gorm:"not null"`
	Cellphone    string `gorm:"not null"`
}
type ContactUserEndereco struct {
	ID           uint   `gorm:"primary_key;autoincrement"`
	Idcoopermapp string `gorm:"not null"`
	Documento    string `gorm:"not null"`
	Cep          string `gorm:"not null"`
	Cidade       string `gorm:"not null"`
	Uf           string `gorm:"not null"`
	Endereco     string `gorm:"not null"`
	Complemento  string
}
type ContactUserVeiculo struct {
	ID             uint   `gorm:"primary_key;autoincrement"`
	Idcoopermapp   string `gorm:"not null"`
	Documento      string `gorm:"not null"`
	Modalidade     int
	Numerocnh      string
	Categoriacnh   string
	Validadecnh    time.Time
	Compareceu     bool
	Uniformizado   bool
	Carroplaca     string
	Carrotipo      int
	Carromodelo    int
	Carromarca     int
	Carroano       string
	Cor            string
	Carga          bool
	Capacidade     float32
	Adesivado      bool
	Dataadesivado  time.Time
	Vistoriado     bool
	Datavistoriado time.Time
}
type CreateContactUserVeiculoInput struct {
	Idcoopermapp   string    `json:"idcoopermapp" binding:"required"`
	Cpf            string    `json:"cpf" binding:"required"`
	Modalidade     int       `json:"modalidade"`
	Numerocnh      string    `json:"numerocnh" `
	Categoriacnh   string    `json:"categoriacnh"`
	Validadecnh    time.Time `json:"validadecnh"`
	Compareceu     bool      `json:"compareceu"`
	Uniformizado   bool      `json:"uniformizado"`
	Carroplaca     string    `json:"carroplaca"`
	Carrotipo      int       `json:"carrotipo"`
	Carromodelo    int       `json:"carromodelo"`
	Carromarca     int       `json:"carromarca"`
	Carroano       string    `json:"carroano"`
	Cor            string    `json:"cor"`
	Carga          bool      `json:"carga"`
	Capacidade     float32   `json:"capacidade"`
	Adesivado      bool      `json:"adesivado"`
	Dataadesivado  time.Time `json:"dataadesivado"`
	Vistoriado     bool      `json:"vistoriado"`
	Datavistoriado time.Time `json:"datavistoriado"`
}

// CreateContactInput : struct for create contact post request
type CreateContactUserInput struct {
	Idcoopermapp string `json:"idcoopermapp" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Cpf          string `json:"cpf" binding:"required"`
	Cellphone    string `json:"cellphone" binding:"required"`
}
type CreateContactUserEnderecoInput struct {
	Idcoopermapp string `json:"idcoopermapp" binding:"required"`
	Cpf          string `json:"cpf" binding:"required"`
	Cep          string `json:"cep" binding:"required"`
	Cidade       string `json:"cidade" binding:"required"`
	Uf           string `json:"uf" binding:"required"`
	Endereco     string `json:"endereco" binding:"required"`
	Complemento  string `json:"complemento"`
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
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)

	tableName := "contato"
	item, err := attributevalue.MarshalMap(&ContactUser{
		Idcoopermapp: input.Idcoopermapp,
		Name:         input.Name,
		Email:        input.Email,
		Documento:    input.Cpf,
		Cellphone:    input.Cellphone,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error marshal map": err.Error()})
		return
	}
	putItemInput := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(c.Request.Context(), putItemInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error put item": err.Error()})
		return
	}
	// Create contact sqlite
	user := ContactUser{Idcoopermapp: input.Idcoopermapp, Name: input.Name, Email: input.Email, Documento: input.Cpf}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateUserEndereco(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input CreateContactUserEnderecoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error input endereco": err.Error()})
		return
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contatoendereco"

	item, err := attributevalue.MarshalMap(&ContactUserEndereco{
		Idcoopermapp: input.Idcoopermapp,
		Documento:    input.Cpf,
		Cep:          input.Cep,
		Cidade:       input.Cidade,
		Uf:           input.Uf,
		Endereco:     input.Endereco,
		Complemento:  input.Complemento,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error marshal map endereco": err.Error()})
		return
	}
	putItemInput := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(c.Request.Context(), putItemInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error put item endereco": err.Error()})
		return
	}
	// Create contact sqlite
	endereco := ContactUserEndereco{Idcoopermapp: input.Idcoopermapp, Documento: input.Cpf, Cep: input.Cep, Cidade: input.Cidade, Uf: input.Uf, Endereco: input.Endereco, Complemento: input.Complemento}

	db.Create(&endereco)

	c.JSON(http.StatusOK, gin.H{"data": endereco})
}

func CreateUserVeiculo(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateContactUserVeiculoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error input veiculo": err.Error()})
		return
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contatoveiculo"

	item, err := attributevalue.MarshalMap(&ContactUserVeiculo{
		Idcoopermapp:   input.Idcoopermapp,
		Documento:      input.Cpf,
		Modalidade:     input.Modalidade,
		Numerocnh:      input.Numerocnh,
		Categoriacnh:   input.Categoriacnh,
		Validadecnh:    input.Validadecnh,
		Compareceu:     input.Compareceu,
		Uniformizado:   input.Uniformizado,
		Carroplaca:     input.Carroplaca,
		Carrotipo:      input.Carrotipo,
		Carromodelo:    input.Carromodelo,
		Carromarca:     input.Carromarca,
		Carroano:       input.Carroano,
		Cor:            input.Cor,
		Carga:          input.Carga,
		Capacidade:     input.Capacidade,
		Adesivado:      input.Adesivado,
		Dataadesivado:  input.Dataadesivado,
		Vistoriado:     input.Vistoriado,
		Datavistoriado: input.Datavistoriado,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error marshal map": err.Error()})
		return
	}
	putItemInput := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(c.Request.Context(), putItemInput)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error put item": err.Error()})
		return
	}
	// Create contact sqlite
	veiculo := ContactUserVeiculo{Idcoopermapp: input.Idcoopermapp, Documento: input.Cpf, Modalidade: input.Modalidade, Numerocnh: input.Numerocnh,
		Categoriacnh: input.Categoriacnh, Validadecnh: input.Validadecnh, Compareceu: input.Compareceu,
		Uniformizado: input.Uniformizado, Carroplaca: input.Carroplaca, Carrotipo: input.Carrotipo,
		Carromodelo: input.Carromodelo, Carromarca: input.Carromarca, Carroano: input.Carroano, Cor: input.Cor,
		Carga: input.Carga, Capacidade: input.Capacidade, Adesivado: input.Adesivado, Dataadesivado: input.Dataadesivado,
		Vistoriado: input.Vistoriado, Datavistoriado: input.Datavistoriado}

	db.Create(&veiculo)

	c.JSON(http.StatusOK, gin.H{"data": veiculo})
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
