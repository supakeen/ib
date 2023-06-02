package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

var AccessToken string

func EnsureToken() {
	if len(AccessToken) == 0 {
		RefreshToken()
	}

	url := "https://console.redhat.com/api/image-builder/v1/distributions"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 401 {
		RefreshToken()
	}
}

func RefreshToken() {
	token := os.Getenv("REDHAT_OFFLINE_TOKEN")

	if len(token) == 0 {
		log.Fatal("No `REDHAT_OFFLINE_TOKEN` in environment.")
	}

	data := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {"rhsm-api"},
		"refresh_token": {token},
	}

	url := "https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(data.Encode())))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var refreshResponse RefreshResponse

	err = json.Unmarshal(body, &refreshResponse)

	if err != nil {
		log.Fatal(err)
	}

	AccessToken = refreshResponse.AccessToken
}
