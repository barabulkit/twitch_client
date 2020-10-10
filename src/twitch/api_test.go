package twitch

import (
	"testing"
)

func TestAccessTokenReceived(t *testing.T) {
	token, err := GetToken()

	if err != nil {
		t.Errorf("failed")
	}

	if token.AccessToken == "" {
		t.Errorf("failed")
	}

}

func TestGamesReceiving(t *testing.T) {
	token, err := GetToken()

	var gamesDtoReq = GamesDtoReq{
		Token: token,
		After: "",
	}

	games, err := GetTopGames(gamesDtoReq)

	if err != nil {
		t.Errorf("faield")
	}

	if len(games.Games) != 20 {
		t.Errorf("failed to recieve games from twitch")
	}
}
