package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gameslog/renderer"
	"gameslog/views"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &renderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Basic Route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get something
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", views.Hello("Person"))
	})

	return r

}
