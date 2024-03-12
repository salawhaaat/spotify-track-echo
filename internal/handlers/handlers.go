package handlers

import (
	"fmt"
	"net/http"
	"spotify-track-echo/internal/client"
	"spotify-track-echo/internal/utils"

	"github.com/gin-gonic/gin"
)

func HandleAlbum(ctx *gin.Context) {
	link := ctx.Param("q")
	if link == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'q' parameter"})
		return
	}
	tracks, err := utils.GetAlbum(link, client.Client)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tracks)
}

func HandlePlaylist(ctx *gin.Context) {
	link := ctx.Param("q")
	if link == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'q' parameter"})
		return
	}
	tracks, err := utils.GetPlaylist(link, client.Client)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tracks)
}

func HandleTrack(ctx *gin.Context) {
	link := ctx.Param("q")
	if link == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'q' parameter"})
		return
	}
	track, err := utils.GetTrack(link, client.Client)
	fmt.Println(track, err)
	if err != nil {
		ctx.JSON(http.StatusAccepted, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, track)
}
