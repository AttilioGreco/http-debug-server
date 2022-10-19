package main

import (
	"io/ioutil"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	// Get port from environment variable
	http_port, exists := os.LookupEnv("HTTP_PORT")
	if !exists {
		http_port = "8080"
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(logger.SetLogger())

	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	r.NoRoute(func(c *gin.Context) {
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		url := c.Request.URL
		log.Info().Msgf("requestUrl: %s, body: %s", url, jsonData)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"data": string(jsonData)})
	})
	r.Run(http_port)
}
