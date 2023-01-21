package routers

import (
	"bikefit/controllers"
	"bikefit/utils"
	"github.com/gin-gonic/gin"
	"html/template"
)

func IniciaRoteamento() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formataData":  utils.FormataData,
		"formataFloat": utils.FormataFloat,
		"contador":     utils.Contador,
	})
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")
	r.GET("/", controllers.Index)
	r.POST("/resultado", controllers.Resultados)
	r.GET("/calculosanteriores", controllers.CalculosAnteriores)
	r.POST("/calculosanteriores", controllers.CalculosAnteriores)
	r.GET("/muraldemensagens", controllers.ExibeMural)
	r.POST("/muraldemensagens", controllers.PostaMural)
	r.GET("/links", controllers.Links)
	r.GET("/sobre", controllers.Sobre)
	r.DELETE("/muraldemensagens/:id", controllers.ApagaMensagem)
	r.Run(":5000")
}
