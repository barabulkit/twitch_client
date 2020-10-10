package twitch

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var props = config.LoadConfig()

func GetToken() (Token, error) {
	var baseURL = "https://id.twitch.tv/oauth2/token"
	data := url.Values{}
	for key, value := range props {
		data.Add(key, value)
	}
	data.Add("grant_type", "client_credentials")
	var token Token

	resp, err := http.PostForm(baseURL, data)
	defer resp.Body.Close()

	if err != nil {
		return token, err
	}

	bytesData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("data", string(bytesData))
	err = json.Unmarshal(bytesData, &token)

	if err != nil {
		return token, err
	}

	token.ClientID = data.Get("client_id")

	return token, nil
}

func GetTopGames(gameDto GamesDtoReq) (GamesDtoResp, error) {
	var baseURL = "https://api.twitch.tv/helix/games/top"
	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL, nil)

	req.Header.Add("Authorization", "Bearer "+gameDto.Token.AccessToken)
	req.Header.Add("Client-ID", gameDto.Token.ClientID)
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		return GamesDtoResp{}, err
	}
	bytesData, _ := ioutil.ReadAll(resp.Body)
	var gamesDtoResp = GamesDtoResp{}
	err = json.Unmarshal(bytesData, &gamesDtoResp)
	fmt.Println(gamesDtoResp)

	return gamesDtoResp, nil
}
