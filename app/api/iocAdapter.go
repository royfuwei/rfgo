package appApi

import (
	"github.com/gin-gonic/gin"
	controller "github.com/royfuwei/rfgo/app/api/controllers"
	appUseCase "github.com/royfuwei/rfgo/pkg/modules/app/usecase"
	jwtService "github.com/royfuwei/rfgo/pkg/modules/jwt/service"
)

type IocAdapter struct{}

func NewIocAdapter() *IocAdapter {
	return &IocAdapter{}
}

func (ioc *IocAdapter) Start(e *gin.Engine) {
	/* ioc */
	appUCase := appUseCase.NewAppUseCase()
	jwtSvc := jwtService.NewJwtService("api/jwt/jwt.rsa", "api/jwt/jwt.rsa.pub")

	/* di */
	appDI := controller.AppDI{
		AppUCase: appUCase,
	}
	controller.NewAppController(e, appDI)
	controller.NewJwtController(e, controller.DI{
		JwtService: jwtSvc,
	})
}
