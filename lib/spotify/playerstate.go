package spotify

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rustic-beans/spotify-viewer/ent"
	"github.com/rustic-beans/spotify-viewer/ent/schema/pulid"
	"github.com/rustic-beans/spotify-viewer/utils"
	"go.uber.org/zap"

	spotifyLib "github.com/zmb3/spotify/v2"
)

// Global playerstate variable
var playerState *PlayerState

type PlayerState struct {
	Track     ent.Track // This the track struct from the DB schema
	Progress  int       `json:"progress_ms"` // This is the current progress of the track in ms
	Date_Time time.Time // The time that the struct was updated last
}

func getPlayer(sa *Spotify) echo.HandlerFunc {
	return func(c echo.Context) error {
		player, err := sa.GetPlayerState(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, player)
	}
}

func PlayerStateLoop(sa *Spotify, dbClient *ent.Client) {
	// Sleep for 5 seconds to give the server time to start
	time.Sleep(5 * time.Second)

	//TODO: Add a loop restart if error occours
	ctx := context.Background()

	defer dbClient.Close()

	// This is the initial playerstate
	playerState = &PlayerState{}
	for {
		player, err := sa.GetPlayerState(ctx)
		if err != nil {
			// This will more than likely happen in the case where nothing is playing or authentication
			// fails
			utils.Logger.Error("Error getting playerstate", zap.Error(err))
		}

		// Check if player is not nil and that the player has an item
		if player != nil && player.Item != nil {

			// Create a track from the playerstate
			track := makeTrack(player)

			// This function requires data from the previous loop so it needs to be called before the update to the playerstate
			// This is to check if the track has changed and if so add it to the db or if the track has been replayed
			dbCheckUpdate(dbClient, track, player.Progress, ctx)

			// This function updates the playerstate with the new track and progress
			updatePlayerState(track, player.Progress)
		}

		// For testing to see if the loop is working
		utils.Logger.Debug("Playerstate recieved", zap.Any("player", playerState))
		// Debugging Query to see if the tracks are being added to the db correctly
		// Best to use len since it removes some of the clutter from the log
		tr, err := dbClient.Track.Query().All(ctx)
		if err != nil {
			utils.Logger.Error("Error querying tracks", zap.Error(err))
		}
		utils.Logger.Debug("Tracks", zap.Any("tracks", len(tr)))

		time.Sleep(5 * time.Second)
	}
}

func makeTrack(player *spotifyLib.PlayerState) *ent.Track {
	// Check if playerState is active if not return previous track as fallback

	artists := make([]string, len(player.Item.Artists))
	//TODO: For some reason the genres field is not available
	// From Item.Artists despite the API saying so. This may be connected
	// to a type issue where Item is not locked to being a track object
	// genres := make(map[string]byte)
	// e.g. player.Item.SimpleTrack.Artists[0].Genres

	for i, artist := range player.Item.Artists {
		artists[i] = artist.Name
	}
	track := &ent.Track{
		Name:          player.Item.Name,
		Artists:       artists,
		ArtistsGenres: nil,
		AlbumName:     player.Item.Album.Name,
		AlbumImageURI: player.Item.Album.Images[0].URL,
		DurationMs:    int32(player.Item.Duration),
		URI:           string(player.Item.URI),
	}

	// Return a pointer to a track
	return track

}

// since playerstate is a public variable we can just update the values inside the pointer instead of returning
func updatePlayerState(track *ent.Track, progress int) {
	if track.Name == "" {
		return
	}

	playerState.Track = *track
	playerState.Progress = progress
	playerState.Date_Time = time.Now()
}

func dbCheckUpdate(dbClient *ent.Client, track *ent.Track, progress int, ctx context.Context) {

	// Check if the track has just changed and if so add it to the db
	if playerState.Track.Name != track.Name {
		addTrack(dbClient, track, ctx)
	}

	// Check for replays
	// TODO: Maybe find a better way to do this but works for now
	// Check if last track update duration is more than 50% done and if current progress is less than 05% into the track
	// This is what constitutes as a replay
	if (int(track.DurationMs/50)*100) < playerState.Progress && progress <= int((track.DurationMs/05)*100) {
		addTrack(dbClient, track, ctx)
	}

}

func addTrack(dbClient *ent.Client, track *ent.Track, ctx context.Context) {
	// NOTE: AritstGenres is not implemented yet so it will be nil see makeTrack function for more info
	_, err := dbClient.Track.Create().
		SetName(track.Name).
		SetArtists(track.Artists).
		SetTrackID(pulid.MustNew(track.Name)).
		SetArtistsGenres(track.ArtistsGenres).
		SetAlbumName(track.AlbumName).
		SetAlbumImageURI(track.AlbumImageURI).
		SetDurationMs(track.DurationMs).
		SetURI(track.URI).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		utils.Logger.Error("Error creating track", zap.Error(err))
	}

}