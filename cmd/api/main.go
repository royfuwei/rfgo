package main

import (
	appApi "github.com/royfuwei/rfgo/app/api"
)

func main() {
	forever := make(chan bool)
	go func() {
		apiService := appApi.NewAPIService()
		apiService.Start()
	}()
	<-forever
}
