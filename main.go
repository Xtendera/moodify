package main

import (
	"fmt"
	"log"
	"moodify/api"

	flag "github.com/spf13/pflag"
)

func main() {
	mood := flag.String("mood", "happy", "target mood for your playlist.")
	clientID := flag.String("cid", "", "the spotify application ID.")
	clientSecret := flag.String("secret", "", "the spotify application secret.")
	playlistID := flag.String("playlist", "", "The playlist ID that should be sorted.")

	flag.Parse()

	token, err := api.GetToken(*clientID, *clientSecret)

	if err != nil {
		log.Fatalf("auth: %s", err)
	}

	fmt.Printf("Mood: %s", *mood)

	playlist, err := api.GetPlaylist(token, *playlistID)
	if err != nil {
		log.Fatalf("getPlaylist: %s", err)
	}
	for _, track := range playlist.Tracks.Items {
		fmt.Printf("\nTrack: %s", track.Track.ID)
	}
}
