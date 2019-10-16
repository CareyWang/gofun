package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofun/routers"
)

func main() {
	r := gin.Default()

	routers.Init(r)

	err := r.Run(":65535")
	if err != nil {
		fmt.Println(err)
	}
}
