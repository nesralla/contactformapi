package api

import (
	"context"
	"fmt"
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
	ID              uint      `gorm:"primary_key;autoincrement"`
	Idcoopermapp    string    `gorm:"not null"`
	Name            string    `gorm:"not null"`
	Email           string    `gorm:"not null"`
	Documento       string    `gorm:"not null"`
	Cellphone       string    `gorm:"not null"`
	Rg              string    `gorm:"not null"`
	DataNascimento  time.Time `gorm:"not null"`
	EstadoCivil     string    `gorm:"not null"`
	NamePai         string    `gorm:"not null"`
	NameMae         string    `gorm:"not null"`
	Sexo            string    `gorm:"not null"`
	Pis             string
	TituloEleitor   string
	Beneficio       bool      `gorm:"not null"`
	DataCadastro    time.Time `gorm:"not null"`
	DataAtualizacao time.Time `gorm:"not null"`
}
type ContactUserEndereco struct {
	ID              uint   `gorm:"primary_key;autoincrement"`
	Idcoopermapp    string `gorm:"not null"`
	Documento       string `gorm:"not null"`
	Cep             string `gorm:"not null"`
	Cidade          string `gorm:"not null"`
	Uf              string `gorm:"not null"`
	Endereco        string `gorm:"not null"`
	Complemento     string
	DataCadastro    time.Time `gorm:"not null"`
	DataAtualizacao time.Time `gorm:"not null"`
}
type ContactUserVeiculo struct {
	ID              uint   `gorm:"primary_key;autoincrement"`
	Idcoopermapp    string `gorm:"not null"`
	Documento       string `gorm:"not null"`
	Modalidade      int
	Numerocnh       string
	Categoriacnh    string
	Validadecnh     time.Time
	Compareceu      bool
	Uniformizado    bool
	Carroplaca      string
	Renavam         string
	Chassi          string
	Carrotipo       int
	Carromodelo     int
	Carromarca      int
	Carroano        string
	Cor             string
	Carga           bool
	Capacidade      float32
	Adesivado       bool
	Dataadesivado   time.Time
	Vistoriado      bool
	Datavistoriado  time.Time
	DataCadastro    time.Time `gorm:"not null"`
	DataAtualizacao time.Time `gorm:"not null"`
}
type CreateContactUserVeiculoInput struct {
	Idcoopermapp   string    `json:"idcoopermapp" binding:"required"`
	Cpf            string    `json:"cpf" binding:"required"`
	Modalidade     int       `json:"modalidade,string,omitempty"`
	Numerocnh      string    `json:"numerocnh" `
	Categoriacnh   string    `json:"categoriacnh"`
	Validadecnh    time.Time `json:"validadecnh"`
	Compareceu     bool      `json:"compareceu"`
	Uniformizado   bool      `json:"uniformizado"`
	Carroplaca     string    `json:"carroplaca"`
	Renavam        string    `json:"renavam"`
	Chassi         string    `json:"chassi"`
	Carrotipo      int       `json:"carrotipo,string,omitempty"`
	Carromodelo    int       `json:"carromodelo,string,omitempty"`
	Carromarca     int       `json:"carromarca,string,omitempty"`
	Carroano       string    `json:"carroano"`
	Cor            string    `json:"cor"`
	Carga          bool      `json:"carga"`
	Capacidade     float32   `json:"capacidade,string,omitempty"`
	Adesivado      bool      `json:"adesivado"`
	Dataadesivado  time.Time `json:"dataadesivado"`
	Vistoriado     bool      `json:"vistoriado"`
	Datavistoriado time.Time `json:"datavistoriado"`
}

// CreateContactInput : struct for create contact post request
type CreateContactUserInput struct {
	Idcoopermapp   string    `json:"idcoopermapp" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Cpf            string    `json:"cpf" binding:"required"`
	Cellphone      string    `json:"cellphone" binding:"required"`
	Rg             string    `json:"rg" binding:"required"`
	DataNascimento time.Time `json:"datanascimento" binding:"required"`
	EstadoCivil    string    `json:"estadocivil" binding:"required"`
	NamePai        string    `json:"namepai" binding:"required"`
	NameMae        string    `json:"namemae" binding:"required"`
	Sexo           string    `json:"sexo" binding:"required"`
	Pis            string    `json:"pis" binding:"required"`
	TituloEleitor  string    `json:"tituloeleitor" binding:"required"`
	Beneficio      bool      `json:"beneficio"`
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
func FindContacts(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)

	var contacts []ContactUser
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contato"
	response, err := client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM \"%v\"", tableName)),
	})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error exe statment ": err.Error()})
		return
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &contacts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error marshall lst": err.Error()})
			return
		}
	}

	//db.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}
func FindContactsEnderecoByCpfAndId(c *gin.Context) {
	var contacts []ContactUserEndereco
	cpf := c.Params.ByName("cpf")
	idcoopermapp := c.Params.ByName("idcoopermapp")
	params, err := attributevalue.MarshalList([]interface{}{cpf, idcoopermapp})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error PArams querystring": err.Error()})
		return
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contatoendereco"
	response, err := client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM \"%v\" WHERE Documento=? AND Idcoopermapp=?", tableName)),
		Parameters: params,
	})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error exe statment ": err.Error()})
		return
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &contacts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error marshall find endereco by cpf and id": err.Error()})
			return
		}
	}

	//db.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}
func FindContactsEndereco(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)

	var contacts []ContactUserEndereco
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contatoendereco"
	response, err := client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM \"%v\"", tableName)),
	})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error exe statment ": err.Error()})
		return
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &contacts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error marshall lst": err.Error()})
			return
		}
	}

	//db.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}

func FindContactsVeiculoByCpfAndId(c *gin.Context) {
	var contacts []ContactUserVeiculo
	cpf := c.Params.ByName("cpf")
	idcoopermapp := c.Params.ByName("idcoopermapp")
	params, err := attributevalue.MarshalList([]interface{}{cpf, idcoopermapp})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error PArams querystring": err.Error()})
		return
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contatoveiculo"
	response, err := client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM \"%v\" WHERE Documento=? AND Idcoopermapp=?", tableName)),
		Parameters: params,
	})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error exe statment ": err.Error()})
		return
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &contacts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error marshall find veiculo by cpf and id": err.Error()})
			return
		}
	}

	//db.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
}
func FindContactsVeiculo(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)

	var contacts []ContactUserVeiculo
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contatoveiculo"
	response, err := client.ExecuteStatement(context.TODO(), &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM \"%v\"", tableName)),
	})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error exe statment ": err.Error()})
		return
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &contacts)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error marshall lst veiculo": err.Error()})
			return
		}
	}

	//db.Find(&contacts)

	c.JSON(http.StatusOK, gin.H{"data": contacts})
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
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config aws": err.Error()})
		return
	}
	client := dynamodb.NewFromConfig(cfg)
	tableName := "contato"
	item, err := attributevalue.MarshalMap(&ContactUser{
		Idcoopermapp:    input.Idcoopermapp,
		Name:            input.Name,
		Email:           input.Email,
		Documento:       input.Cpf,
		Cellphone:       input.Cellphone,
		Rg:              input.Rg,
		DataNascimento:  input.DataNascimento,
		EstadoCivil:     input.EstadoCivil,
		NamePai:         input.NamePai,
		NameMae:         input.NameMae,
		Sexo:            input.Sexo,
		Pis:             input.Pis,
		TituloEleitor:   input.TituloEleitor,
		Beneficio:       input.Beneficio,
		DataCadastro:    time.Now(),
		DataAtualizacao: time.Now(),
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
	user := ContactUser{Idcoopermapp: input.Idcoopermapp, Name: input.Name, Email: input.Email, Documento: input.Cpf, Cellphone: input.Cellphone,
		Rg: input.Rg, DataNascimento: input.DataNascimento, EstadoCivil: input.EstadoCivil, NamePai: input.NamePai, NameMae: input.NameMae,
		Sexo: input.Sexo, Pis: input.Pis, TituloEleitor: input.TituloEleitor, Beneficio: input.Beneficio, DataCadastro: time.Now(), DataAtualizacao: time.Now()}

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
		Idcoopermapp:    input.Idcoopermapp,
		Documento:       input.Cpf,
		Cep:             input.Cep,
		Cidade:          input.Cidade,
		Uf:              input.Uf,
		Endereco:        input.Endereco,
		Complemento:     input.Complemento,
		DataCadastro:    time.Now(),
		DataAtualizacao: time.Now(),
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
	endereco := ContactUserEndereco{Idcoopermapp: input.Idcoopermapp, Documento: input.Cpf, Cep: input.Cep, Cidade: input.Cidade, Uf: input.Uf, Endereco: input.Endereco, Complemento: input.Complemento, DataCadastro: time.Now(), DataAtualizacao: time.Now()}

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
		Idcoopermapp:    input.Idcoopermapp,
		Documento:       input.Cpf,
		Modalidade:      input.Modalidade,
		Numerocnh:       input.Numerocnh,
		Categoriacnh:    input.Categoriacnh,
		Validadecnh:     input.Validadecnh,
		Compareceu:      input.Compareceu,
		Uniformizado:    input.Uniformizado,
		Carroplaca:      input.Carroplaca,
		Renavam:         input.Renavam,
		Chassi:          input.Chassi,
		Carrotipo:       input.Carrotipo,
		Carromodelo:     input.Carromodelo,
		Carromarca:      input.Carromarca,
		Carroano:        input.Carroano,
		Cor:             input.Cor,
		Carga:           input.Carga,
		Capacidade:      input.Capacidade,
		Adesivado:       input.Adesivado,
		Dataadesivado:   input.Dataadesivado,
		Vistoriado:      input.Vistoriado,
		Datavistoriado:  input.Datavistoriado,
		DataCadastro:    time.Now(),
		DataAtualizacao: time.Now(),
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
		Uniformizado: input.Uniformizado, Carroplaca: input.Carroplaca, Renavam: input.Renavam, Chassi: input.Chassi, Carrotipo: input.Carrotipo,
		Carromodelo: input.Carromodelo, Carromarca: input.Carromarca, Carroano: input.Carroano, Cor: input.Cor,
		Carga: input.Carga, Capacidade: input.Capacidade, Adesivado: input.Adesivado, Dataadesivado: input.Dataadesivado,
		Vistoriado: input.Vistoriado, Datavistoriado: input.Datavistoriado, DataCadastro: time.Now(), DataAtualizacao: time.Now()}

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
