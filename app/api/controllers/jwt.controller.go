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
	router.POST("/decode", handler.JwtDecode)
	router.POST("/verify", handler.JwtVerify)
}

// @Summary Sign jwt token
// @Description	Sign jwt token
// @Accept  json
// @Produce  json
// @Param default body domain.ReqJwtSign true "jwt sign"
// @Success 200 {object} domain.TokenClaimsDTO	"ok"
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

// @Summary Decode jwt token
// @Description	Decode jwt token
// @Accept  json
// @Produce  json
// @Param default body domain.ReqJwtToken true "json web token"
// @Success 200 {object} domain.TokenClaimsDTO	"ok"
// @Router /jwt/decode [post]
func (h *jwtHandler) JwtDecode(c *gin.Context) {
	var req domain.ReqJwtToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := req.Token
	claims, err := h.DI.JwtService.JwtDecode(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, claims)
}

// @Summary Verify jwt token
// @Description	Verify jwt token
// @Accept  json
// @Produce  json
// @Param default body domain.ReqJwtToken true "json web token"
// @Success 200 {object} domain.TokenClaimsDTO	"ok"
// @Router /jwt/verify [post]
func (h *jwtHandler) JwtVerify(c *gin.Context) {
	var req domain.ReqJwtToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := req.Token
	claims, err := h.DI.JwtService.JwtVerify(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, claims)
}
