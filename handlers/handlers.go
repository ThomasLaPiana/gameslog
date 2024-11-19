package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gameslog/database"
	"gameslog/renderer"
	"gameslog/views"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &renderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"detail": "healthy",
		})
	})

	// Index
	r.GET("/", func(c *gin.Context) {
		database := database.NewDatabase()
		games := database.GetAllGames()
		c.HTML(http.StatusOK, "", views.Index(games))
	})

	return r

}
