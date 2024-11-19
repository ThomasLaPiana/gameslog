package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running Webserver!")

	// Create the Router and set the HTML Renderer
	r := gin.Default()
	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Basic Route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get something
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", hello("Person"))
	})

	r.Run()
}
