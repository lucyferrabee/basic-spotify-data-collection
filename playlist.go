package main

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type playlist struct {
	id          string
	name        string
	description string
}

func getById(id string) *spotify.FullPlaylist {

	authConfig := &clientcredentials.Config{
		ClientID:     "4779b9533e004287b6536fd8c5325adf",
		ClientSecret: "8dbf0a481f8b4de0901fbc9661f6036c",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieving access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	playlistID := spotify.ID(id)
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("error retrieving playlist data: %v", err)
	}

	return playlist
}

func createObject(spotifyPlaylist *spotify.FullPlaylist) playlist {
	p := playlist{
		name:        spotifyPlaylist.Name,
		id:          string(spotifyPlaylist.ID),
		description: spotifyPlaylist.Description,
	}

	return p
}
