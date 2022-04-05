package main

import "github.com/SabatiniFederico/mercadolibre-quasar-challenge/router"

func main() {

	server := router.GetRouter()
	server.Run(":8080")
}
