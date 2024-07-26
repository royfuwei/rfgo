package main

import (
	ginrest "github.com/royfuwei/rfgo/cmd/api/infra"
)

func main() {
	forever := make(chan bool)
	go func() {
		apiService := ginrest.NewAPIService()
		apiService.Start()
	}()
	<-forever
}
