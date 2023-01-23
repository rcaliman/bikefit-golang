package controllers

import (
	"bikefit/databases"
	"bikefit/models"
	"bikefit/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "views/index.html", nil)
}

func Resultados(c *gin.Context) {
	cavalo := utils.StringParaFloat(c.PostForm("cavalo"))
	esterno := utils.StringParaFloat(c.PostForm("esterno"))
	braco := utils.StringParaFloat(c.PostForm("braco"))
	email := c.PostForm("email")

	var medidas models.Medidas
	medidas = models.Medidas{
		Cavalo:         cavalo,
		Esterno:        esterno,
		Braco:          braco,
		Email:          email,
		Tronco:         medidas.CalculaTronco(esterno, cavalo),
		QuadroSpeed:    medidas.CalculaSpeed(cavalo),
		QuadroMTB:      medidas.CalculaMTB(cavalo),
		AlturaSelim:    medidas.CalculaSelim(cavalo),
		TopTubeEfetivo: medidas.CalculaTopTube(esterno, cavalo, braco),
	}
	databases.DB.Create(&medidas)
	c.HTML(http.StatusOK, "views/resultado.html", gin.H{
		"resultado": medidas,
	})
}

func CalculosAnteriores(c *gin.Context) {
	email := c.PostForm("email")
	var calculos []models.Medidas
	if email != "" {
		databases.DB.Where("email = ?", email).Find(&calculos)
	}
	c.HTML(http.StatusOK, "views/calculosanteriores.html", gin.H{
		"calculos": calculos,
	})
}

func Links(c *gin.Context) {
	c.HTML(http.StatusOK, "views/links.html", nil)
}

func Sobre(c *gin.Context) {
	c.HTML(http.StatusOK, "views/sobre.html", nil)
}
