package main

import (
	"fmt"

	"github.com/rustic-beans/spotify-viewer/internal/utils"
	"go.uber.org/zap"

	"github.com/rustic-beans/spotify-viewer/internal/infrastructure/database"
	"github.com/rustic-beans/spotify-viewer/internal/infrastructure/graphql"
	httpLib "github.com/rustic-beans/spotify-viewer/internal/infrastructure/http"
	"github.com/rustic-beans/spotify-viewer/internal/spotify"
	spotifyLib "github.com/zmb3/spotify/v2"
)

func main() {
	// Logger
	defer func() {
		_ = utils.Logger.Sync()
	}()

	// Config
	config, err := utils.ReadConfig()
	if err != nil {
		utils.Logger.Fatal("failed reading config", zap.Error(err))
	}

	dbClient := database.Connect(config)

	spotifyClient := spotify.NewSpotify(config)

	playerStateWebsocketHandler := httpLib.NewWebsocketHandler[*spotifyLib.PlayerState]()

	graphqlServer := graphql.NewServer(dbClient, spotifyClient, playerStateWebsocketHandler)
	e := httpLib.NewServer(graphqlServer)

	spotifyClient.SetupRoutes(e)
	spotifyClient.Authenticate()

	go spotify.PlayerStateLoop(spotifyClient, dbClient, playerStateWebsocketHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)))
}
