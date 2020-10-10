package twitch

//Token - Twitch auth token
type Token struct {
	ClientID    string
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

//Game - Struct for twitch game response
type Game struct {
	ID        int    `json:"id`
	Name      string `json:"name`
	BoxArtURL string `json:"box_art_url"`
}

//Pagination - Struct for twitch pagination response
type Pagination struct {
	Cursor string `json:"cursor"`
}

type GamesDtoResp struct {
	Games      []Game     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

//GameDtoReq - Struct for twitch games request
// after - cursor position for pagination
type GamesDtoReq struct {
	Token Token
	After string
}
