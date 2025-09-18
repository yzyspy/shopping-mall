package main

import (
	"shopping_mall/models"
	"shopping_mall/router"
)

func main() {
	models.NewDB()

	r := router.App()
	r.Run(":8081")

	//my_shopping_mall
	//admin
	//root
	//10.92.160.243
}
