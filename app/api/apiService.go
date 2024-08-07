package appApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	_ "github.com/royfuwei/rfgo/api/swag" // docs is generated by Swag CLI, you have to import it.
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"
)

type APIService struct{}

func NewAPIService() *APIService {
	return &APIService{}
}

func (api *APIService) Start() {
	r := gin.New()

	// Global middleware
	r.Use((gin.Logger()))
	r.Use(gin.Recovery())

	gin.ForceConsoleColor()

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	iocAdapter := NewIocAdapter()
	iocAdapter.Start(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	glog.Info("Server starting")
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			glog.Info("Server closed under request")
		} else {
			glog.Fatalf("Server closed unexpect: %v", err)
		}
	}
	glog.Info("Server exiting")
}
