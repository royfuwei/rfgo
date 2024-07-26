package appApi

import (
	"github.com/gin-gonic/gin"
	controller "github.com/royfuwei/rfgo/app/api/controllers"
	appUseCase "github.com/royfuwei/rfgo/pkg/modules/app/usecase"
)

type IocAdapter struct{}

func NewIocAdapter() *IocAdapter {
	return &IocAdapter{}
}

func (ioc *IocAdapter) Start(e *gin.Engine) {
	/* ioc */
	appUCase := appUseCase.NewAppUseCase()

	/* di */
	appDI := controller.AppDI{
		AppUCase: appUCase,
	}
	controller.NewAppController(e, appDI)
}
