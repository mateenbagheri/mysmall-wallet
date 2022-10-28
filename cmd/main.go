package main

import (
	"github.com/mateenbagheri/mysmall-wallet/routes"
	"github.com/mateenbagheri/mysmall-wallet/scripts"
)

func main() {
	go scripts.DailySumScheduler()
	router := routes.SetUpRouter()
	router.Run(":8080")
}
