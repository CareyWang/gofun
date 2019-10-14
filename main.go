package main

import (
	"fmt"
	"github.com/CareyWang/gofun/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.Init(r)

	err := r.Run(":65535")
	if err != nil {
		fmt.Println(err)
	}
}
