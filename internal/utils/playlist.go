package utils

import (
	"fmt"

	"github.com/zmb3/spotify"
)

func GetPlaylist(url string, client *spotify.Client) (tracks []string, err error) {
	playlist, err := client.GetPlaylist(spotify.ID(url))
	if err != nil {
		return nil, err
	}
	for _, v := range playlist.Tracks.Tracks {
		tracks = append(tracks, v.Track.Artists[0].Name+
			" "+v.Track.Name)
	}
	return tracks, nil
}

func GetAlbum(url string, client *spotify.Client) ([]string, error) {
	var tracks []string
	playlist, err := client.GetAlbum(spotify.ID(url))
	fmt.Println(playlist)
	if err != nil {
		return nil, err
	}
	for _, v := range playlist.Tracks.Tracks {
		tracks = append(tracks, v.Artists[0].Name+
			" "+v.Name)
	}
	return tracks, nil
}

func GetTrack(url string, client *spotify.Client) (track string, err error) {
	t, err := client.GetTrack(spotify.ID(url))
	if err != nil {
		return "", err
	}
	track = t.Artists[0].Name + " " + t.Name
	return track, nil
}
