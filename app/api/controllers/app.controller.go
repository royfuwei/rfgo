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
	e.GET("/", handler.GetApp)
	router := e.Group("/app")
	router.GET("/", handler.GetApp)
}

// @Summary Get Api app name
// @Description
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.App	"ok"
// @Router /app [get]
func (h *appHandler) GetApp(c *gin.Context) {
	data := h.DI.AppUCase.GetApp()
	c.JSON(http.StatusOK, data)
}
