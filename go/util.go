package swagger

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken() (string, error) {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tokenUrl := os.Getenv("TOKEN_URL")
	fmt.Println(clientId, clientSecret, tokenUrl)
	header := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret))
	r, e := http.NewRequest("POST", tokenUrl, bytes.NewReader([]byte("grant_type=client_credentials")))

	if e != nil {
		fmt.Println(e)
		return "", e
	}
	r.Header.Add("Authorization", header)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(r)

	if err != nil {
		fmt.Println(err)
		return "", e
	}

	accessToken := &TokenResponse{}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return "", e
	}

	json.Unmarshal(body, accessToken)

	return accessToken.AccessToken, nil
}
