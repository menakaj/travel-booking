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

func GetToken() string {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	tokenUrl := os.Getenv("TOKEN_URL")

	fmt.Println(clientId)
	fmt.Println(clientSecret)
	fmt.Println(tokenUrl)

	header := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret))
	r, e := http.NewRequest("POST", tokenUrl, bytes.NewReader([]byte("grant_type=client_credentials")))

	if e != nil {
		fmt.Println(e)
		return ""
	}
	r.Header.Add("Authorization", header)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(r)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	accessToken := &TokenResponse{}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	json.Unmarshal(body, accessToken)

	return accessToken.AccessToken
}
