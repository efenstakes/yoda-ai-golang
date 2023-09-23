package main

import (
	"log"

	"github.com/efenstakes/yoda-ai/prompt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func main() {

	// create default server
	r := gin.Default()

	// index route
	r.GET("/", prompt.Index)

	// prompt route
	r.GET("/prompt", prompt.Prompt)

	// start server
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
