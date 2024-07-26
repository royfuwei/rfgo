package usecase

import "github.com/royfuwei/rfgo/pkg/domain"

type appUseCase struct{}

func NewAppUseCase() domain.AppUseCase {
	return &appUseCase{}
}

func (ucase *appUseCase) GetApp() *domain.App {
	return &domain.App{
		AppName: "rfgo/app",
	}
}
