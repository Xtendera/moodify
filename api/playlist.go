package api

import (
	"fmt"
	"moodify/utils"
)

// Playlist This is not extensive, but I don't need every property.
type Playlist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Tracks struct {
		Items []struct {
			Track struct {
				ID string `json:"id"`
			} `json:"track"`
		} `json:"items"`
	} `json:"tracks"`
}

// GetPlaylist The playlist parameter should be the playlist ID.
func GetPlaylist(token string, playlist string) (*Playlist, error) {
	var playlistObj Playlist
	err := utils.SendURLEncoded("GET", fmt.Sprintf("https://api.spotify.com/v1/playlists/%s", playlist), nil, &playlistObj, &token)
	if err != nil {
		return nil, fmt.Errorf("error fetching playlist: %w", err)
	}
	return &playlistObj, nil
}
