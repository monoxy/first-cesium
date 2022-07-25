package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("frontend/*.html")
	r.Static("/Cesium", "./frontend/Cesium")
	r.Static("/scripts", "./frontend/scripts")
	r.Static("/styles", "./frontend/styles")
	r.Handle("GET", "/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.Run(":9001")
}
