package main

import (
	"github.com/gautamarora04/go-sms-verify-yt/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// initailize config
	app := api.Config{Router: router}

	// routes
	app.Routes()

	router.Run(":8000")
}
