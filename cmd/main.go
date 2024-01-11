package main

import (
	"net/http"

	"github.com/eu-micaeu/IsaSemiJoias-WebSite/middlewares"
	"github.com/gin-gonic/gin"
	
)

func main() {

	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	r.LoadHTMLGlob("./views/*.html")

	r.GET("/index", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", nil)

	})

	r.GET("/semijoia", func(c *gin.Context) {

		c.HTML(http.StatusOK, "semijoia.html", nil)

	})

	r.GET("/pagamento", func(c *gin.Context) {

		c.HTML(http.StatusOK, "pagamento.html", nil)

	})

	r.Static("./static", "./static")

	r.Run()
}