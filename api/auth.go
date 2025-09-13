package api

import (
	"errors"
	"fmt"
	"moodify/utils"
	"net/url"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint   `json:"expires_in"`
}

func GetToken(cid string, secret string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", cid)
	data.Set("client_secret", secret)

	var jsonResp AuthResponse
	err := utils.SendURLEncoded("POST", "https://accounts.spotify.com/api/token", data, &jsonResp, nil)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	if jsonResp.TokenType != "Bearer" {
		return "", errors.New("unsupported authentication type")
	}
	return jsonResp.AccessToken, nil
}
