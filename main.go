package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/navidrome/navidrome/plugins/pdk/go/host"
	"github.com/navidrome/navidrome/plugins/pdk/go/metadata"
	"github.com/navidrome/navidrome/plugins/pdk/go/pdk"
	"github.com/navidrome/navidrome/plugins/pdk/go/types"
)

const configAPIUrl = "apiUrl"
const temperatureTrackSimilarity = "temperatureTrackSimilarity"
const temperatureAlbumSimilarity = "temperatureAlbumSimilarity"
const temperatureArtistSimilarity = "temperatureArtistSimilarity"

type TrackResponse struct {
	Path     string  `json:"path"`
	SongId   string  `json:"songId"`
	AlbumId  string  `json:"albumId"`
	ArtistId string  `json:"artistId"`
	Score    float64 `json:"score"`
}

type clapclapPlugin struct{}

const LOG_PREFIX = "[clapclap]"

func logInfo(text string) {
	pdk.Log(pdk.LogInfo, fmt.Sprintf("%s %s", LOG_PREFIX, text))
}

func logError(text error) {
	pdk.Log(pdk.LogError, fmt.Sprintf("%s %s", LOG_PREFIX, text.Error()))
}

func init() {
	metadata.Register(&clapclapPlugin{})
	logInfo("Plugin Loaded")
}

func (p *clapclapPlugin) GetSimilarSongsByTrack(input metadata.SimilarSongsByTrackRequest) (*metadata.SimilarSongsResponse, error) {
	logInfo(fmt.Sprintf("GetSimilarSongsByTrack (ID: %s, Song: %s, Artist: %s)", input.ID, input.Name, input.Artist))
	temperatureStr, ok := pdk.GetConfig(temperatureTrackSimilarity)
	if !ok {
		return nil, fmt.Errorf("No " + temperatureTrackSimilarity + " set")
	}
	temperature, _ := strconv.ParseFloat(temperatureStr, 64)
	return p.GetSimilarSongs(input.ID, "", "", input.Count, temperature)
}

func (p *clapclapPlugin) GetSimilarSongsByAlbum(input metadata.SimilarSongsByAlbumRequest) (*metadata.SimilarSongsResponse, error) {
	logInfo(fmt.Sprintf("GetSimilarSongsByAlbum (ID: %s, Album: %s, Artist: %s)", input.ID, input.Name, input.Artist))
	temperatureStr, ok := pdk.GetConfig(temperatureAlbumSimilarity)
	if !ok {
		return nil, fmt.Errorf("No " + temperatureAlbumSimilarity + " set")
	}
	temperature, _ := strconv.ParseFloat(temperatureStr, 64)
	return p.GetSimilarSongs("", input.ID, "", input.Count, temperature)
}

func (p *clapclapPlugin) GetSimilarSongsByArtist(input metadata.SimilarSongsByArtistRequest) (*metadata.SimilarSongsResponse, error) {
	logInfo(fmt.Sprintf("GetSimilarSongsByArtist (ID: %s, Artist: %s)", input.ID, input.Name))
	temperatureStr, ok := pdk.GetConfig(temperatureArtistSimilarity)
	if !ok {
		return nil, fmt.Errorf("No " + temperatureArtistSimilarity + " set")
	}
	temperature, _ := strconv.ParseFloat(temperatureStr, 64)
	return p.GetSimilarSongs("", "", input.ID, input.Count, temperature)
}

func (p *clapclapPlugin) GetSimilarSongs(songId string, albumId string, artistId string, count int32, temperature float64) (*metadata.SimilarSongsResponse, error) {
	tracks, err := p.queryAPI(songId, albumId, artistId, count, temperature)
	if err != nil {
		logError(err)
		return nil, err
	}

	// Convert to Navidrome SongRef format preserving order
	songs := make([]types.SongRef, 0, len(tracks))
	for _, track := range tracks {
		songs = append(songs, types.SongRef{
			ID: track.SongId,
		})
	}

	logInfo(fmt.Sprintf("%d songs found", len(songs)))

	return &metadata.SimilarSongsResponse{Songs: songs}, nil
}

func (p *clapclapPlugin) queryAPI(songId string, albumId string, artistId string, count int32, temperature float64) ([]TrackResponse, error) {
	apiBaseURL, ok := pdk.GetConfig(configAPIUrl)

	if !ok {
		return nil, fmt.Errorf("No " + configAPIUrl + " set")
	}

	params := url.Values{}
	params.Set("limit", fmt.Sprint(count))
	params.Set("temperature", fmt.Sprint(temperature))
	if songId != "" {
		params.Set("songId", songId)
	}
	if albumId != "" {
		params.Set("albumId", albumId)
	}
	if artistId != "" {
		params.Set("artistId", artistId)
	}

	apiURL := fmt.Sprintf("%s/similar_songs?%s", apiBaseURL, params.Encode())
	logInfo(fmt.Sprintf("query API: %s", apiURL))

	resp, err := host.HTTPSend(host.HTTPRequest{
		Method: "GET",
		URL:    apiURL,
	})

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("response status %d, body: %s", resp.StatusCode, resp.Body)
	} else {
		logInfo(fmt.Sprintf("response status: %d", resp.StatusCode))
	}

	var tracks []TrackResponse
	if err := json.Unmarshal(resp.Body, &tracks); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return tracks, nil
}

func main() {}
