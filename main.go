package main

import (
	"DurianHamsterDen/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)
	_ = r.Run(":8180")
}
