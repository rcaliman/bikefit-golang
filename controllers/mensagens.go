package controllers

import (
	"bikefit/databases"
	"bikefit/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ExibeMural(c *gin.Context) {
	var mensagens []models.Mensagem
	databases.DB.Order("created_at DESC, id DESC").Find(&mensagens)
	c.HTML(http.StatusOK, "views/mural.html", gin.H{
		"mensagens": mensagens,
	})
}

func PostaMural(c *gin.Context) {
	nome := c.PostForm("nome")
	email := c.PostForm("email")
	texto := c.PostForm("mensagem")

	mensagem := models.Mensagem{
		Nome:  nome,
		Email: email,
		Texto: texto,
	}
	if !strings.Contains(texto, "<a") {
		databases.DB.Create(&mensagem)
	}
	var mensagens []models.Mensagem
	databases.DB.Order("created_at DESC, id DESC").Find(&mensagens)

	c.HTML(http.StatusOK, "views/mural.html", gin.H{
		"mensagens": mensagens,
	})
}
