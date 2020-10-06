package twitch

import (
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "https://id.twitch.tv/oauth2/token"

var props = config.LoadConfig()

func GetToken() (Token, error) {
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

	return token, nil
}
