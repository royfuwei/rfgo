package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/royfuwei/rfgo/pkg/domain"
)

type jwtHandler struct {
	DI
	JwtController
}

type DI struct {
	JwtService domain.JwtService
}

type JwtController interface {
	JwtSign(c *gin.Context)
}

func NewJwtController(e *gin.Engine, di DI) {
	handler := &jwtHandler{
		DI: di,
	}
	router := e.Group("/jwt")
	router.POST("/sign", handler.JwtSign)
}

// @Summary Sign jwt token
// @Description	Sign jwt token
// @Accept  json
// @Produce  json
// @Param default body domain.ReqJwtSign true "account login 內容"
// @Success 200 {object} domain.TokenClaims	"ok"
// @Router /jwt/sign [post]
func (h *jwtHandler) JwtSign(c *gin.Context) {
	var req domain.ReqJwtSign
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	expiresAt, token, err := h.DI.JwtService.JwtSign(time.Duration(time.Duration.Seconds(60*60)), req.Uid, nil)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"expiresAt": expiresAt, "token": token})
}
