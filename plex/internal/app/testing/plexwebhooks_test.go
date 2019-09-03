package model_testing

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/dabiggm0e/plextrakt/plex/internal/app/model"
)

var movie = model.MediaTypeMovie
var show = model.MediaTypeShow
var scrobble = model.EventScrobble

type event struct {
	user          string
	eventtype     string
	mediatype     string
	title         string
	season        int
	episodeNumber int
	episodeTitle  string
}

func TestMarshalingEvents(t *testing.T) {

	eventTests := []struct {
		filename string
		want     event
	}{
		{filename: "samples/movie-scrobble.json",
			want: event{user: "dabiggmoe", eventtype: "media.scrobble",
				mediatype: "movie", title: "Angels & Demons"}},

		{filename: "samples/anime-scrobble.json",
			want: event{user: "dabiggmoe", eventtype: "media.scrobble",
				mediatype: "show", title: "Gintama", episodeNumber: 1,
				season: 1, episodeTitle: "Eat Something Sour When You're Tired!"}},

		{filename: "samples/anime-scrobble.json",
			want: event{user: "dabiggmoe", eventtype: "media.scrobble",
				mediatype: "show", title: "Gintama", episodeNumber: 1,
				season: 1, episodeTitle: "Eat Something Sour When You're Tired!"}},

		{filename: "samples/show-scrobble2.json",
			want: event{user: "dabiggmoe", eventtype: "media.scrobble",
				mediatype: "show", title: "Peaky Blinders", episodeNumber: 1,
				season: 5, episodeTitle: "Black Tuesday"}},
	}

	for _, testcase := range eventTests {
		payload, err := loadEventFromFile(testcase.filename)
		if err != nil {
			t.Errorf("Error loading file: %v", err)
			continue
		}

		got, err := model.NewEventFromJson(payload)
		if err != nil {
			t.Errorf("Error loading json: %v", err)
			continue
		}

		switch got.GetMediaType() {
		case movie:
			assertEqualString(t, got.GetAccountUser(), testcase.want.user)
			assertEqualString(t, got.GetEvent(), testcase.want.eventtype)
			assertEqualString(t, got.GetMediaType(), testcase.want.mediatype)
			assertEqualString(t, got.GetMovieTitle(), testcase.want.title)

		case show:
			assertEqualString(t, got.GetAccountUser(), testcase.want.user)
			assertEqualString(t, got.GetEvent(), testcase.want.eventtype)
			assertEqualString(t, got.GetMediaType(), testcase.want.mediatype)
			assertEqualString(t, got.GetShowTitle(), testcase.want.title)
			assertEqualString(t, got.GetEpisodeTitle(), testcase.want.episodeTitle)
			assertEqualInt(t, got.GetEpisodeNo(), testcase.want.episodeNumber)
			assertEqualInt(t, got.GetSeason(), testcase.want.season)
		}

	}
}

func loadEventFromFile(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Printf("Error opening file %v: %v", filename, err)
		return "", err
	}

	return string(file), nil
}

func assertEqualString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got '%v' want '%v'", got, want)
	}
}

func assertEqualInt(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
