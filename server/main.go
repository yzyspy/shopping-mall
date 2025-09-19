package main

import (
	"shopping_mall/models"
	"shopping_mall/router"
)

func main() {
	models.NewDB()

	r := router.App()
	
	r.Run(":8081")
}
