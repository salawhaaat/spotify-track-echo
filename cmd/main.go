package main

import (
	"io"
	"os"
	"spotify-track-echo/internal/client"
	"spotify-track-echo/internal/config"
	"spotify-track-echo/internal/handlers"
	"spotify-track-echo/pkg/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	cfg := config.New()
	client.New(cfg)

}

func main() {
	r := gin.Default()
	f, _ := os.Create("logs/app.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.LoggerWithFormatter(logger.FormatLogs))
	r.GET("/playlist/:q", handlers.HandlePlaylist)
	r.GET("/album/:q", handlers.HandleAlbum)
	r.GET("/track/:q", handlers.HandleTrack)
	r.Run()
}
