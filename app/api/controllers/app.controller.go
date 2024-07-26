package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/royfuwei/rfgo/pkg/domain"
)

type appHandler struct {
	DI AppDI
	AppController
}

type AppDI struct {
	AppUCase domain.AppUseCase
}

type AppController interface {
	GetApp(c *gin.Context)
}

func NewAppController(e *gin.Engine, di AppDI) {
	handler := &appHandler{
		DI: di,
	}
	router := e.Group("/")
	router.GET("", handler.GetApp)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (h *appHandler) GetApp(c *gin.Context) {
	data := h.DI.AppUCase.GetApp()
	c.JSON(http.StatusOK, data)
}
