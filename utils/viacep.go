package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func FindAdressByZipCode(c *gin.Context) {
	var viaCep ViaCep
	zipcode := c.Params.ByName("cep")
	req, err := http.Get("https://viacep.com.br/ws/" + zipcode + "/json/")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error config ws cep": err.Error()})
		return
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error io read ": err.Error()})
		return
	}
	err = json.Unmarshal(res, &viaCep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error unarshal ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": viaCep})
}
