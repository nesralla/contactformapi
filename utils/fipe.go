package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FipeOutput struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
type MarcasInput struct {
	CodigoTabelaReferencia int `json:"codigoTabelaReferencia"`
	CodigoTipoVeiculo      int `json:"codigoTipoVeiculo"`
}
type ModelosInput struct {
	CodigoTabelaReferencia int `json:"codigoTabelaReferencia"`
	CodigoTipoVeiculo      int `json:"codigoTipoVeiculo"`
	CodigoMarca            int `json:"codigoMarca"`
}

func GetMarca(c *gin.Context) {
	var fipeOutputs []FipeOutput
	codigo := c.Params.ByName("codigo")
	tipo, err := strconv.Atoi(c.Params.ByName("tipo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error convert tipo int": err.Error()})
		return
	}
	body, _ := json.Marshal(map[string]int{
		"codigoTabelaReferencia": 231,
		"codigoTipoVeiculo":      tipo,
	})
	payload := bytes.NewBuffer(body)
	req, err := http.Post("http://veiculos.fipe.org.br/api/veiculos/ConsultarMarcas", "application/json", payload)
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
	err = json.Unmarshal(res, &fipeOutputs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error unarshal ": err.Error()})
		return
	}
	out := filter(fipeOutputs, codigo)
	c.JSON(http.StatusOK, gin.H{"data": out})
}
func filter(in []FipeOutput, codigo string) []FipeOutput {
	var out []FipeOutput
	for _, each := range in {
		if each.Value == codigo {
			out = append(out, each)
		}
	}
	return out
}
func GetModelo(c *gin.Context) {
	var fipeOutputs []FipeOutput
	codigo := c.Params.ByName("codigo")
	tipo, err := strconv.Atoi(c.Params.ByName("tipo"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error convert tipo int": err.Error()})
		return
	}
	marca, err := strconv.Atoi(c.Params.ByName("marca"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error convert tipo int": err.Error()})
		return
	}
	body, _ := json.Marshal(map[string]int{
		"codigoTabelaReferencia": 301,
		"codigoTipoVeiculo":      tipo,
		"codigoMarca":            marca,
	})
	payload := bytes.NewBuffer(body)
	req, err := http.Post("http://veiculos.fipe.org.br/api/veiculos/ConsultarModelos", "application/json", payload)
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
	err = json.Unmarshal(res, &fipeOutputs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error unarshal ": err.Error()})
		return
	}
	out := filter(fipeOutputs, codigo)
	c.JSON(http.StatusOK, gin.H{"data": out})
}
