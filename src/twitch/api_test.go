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
