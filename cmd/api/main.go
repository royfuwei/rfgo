package main

import (
	appApi "github.com/royfuwei/rfgo/app/api"
)

// @title rfgo open API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath
func main() {
	forever := make(chan bool)
	go func() {
		apiService := appApi.NewAPIService()
		apiService.Start()
	}()
	<-forever
}
