package appApi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
