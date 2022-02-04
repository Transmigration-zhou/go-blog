package main

import (
	"go-blog/model"
	"go-blog/router"
	"go-blog/utils"
)

func main() {
	utils.InitConfig()
	model.InitDB()
	router.InitRouters()
}
