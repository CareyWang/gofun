package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gofun/routes"
)

func main() {
	r := gin.Default()

	routes.Init(r)

	err := r.Run(":65535")
	if err != nil {
		fmt.Println(err)
	}
}
