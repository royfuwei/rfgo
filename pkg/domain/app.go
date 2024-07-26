package domain

type App struct {
	AppName string `json:"app,omitempty"`
}

type AppUseCase interface {
	GetApp() *App
}
