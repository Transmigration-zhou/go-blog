package main

import (
	"go-blog/model"
	"go-blog/utils"
)

func main() {
	utils.InitConfig()
	model.InitDB()
	//routers.InitRouters()
}
