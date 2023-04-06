package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	API_KEY    = "justpp"
	API_SECRET = "tour/blog"
)

type AccessToken struct {
	Token string `json:"token"`
}

type ApiIntFc interface {
	httpGet(ctx context.Context, path string) ([]byte, error)
	getAccessToken(ctx context.Context) (string, error)
}

func NewAPI(url string) *API {
	return &API{url}
}

type API struct {
	URL string
}

func (a *API) httpGet(ctx context.Context, path string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", a.URL, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (a *API) getAccessToken(ctx context.Context) (string, error) {
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?app_key=%s&app_secret=%s", "auth", API_KEY, API_SECRET))
	if err != nil {
		return "", err
	}
	var accessToken AccessToken
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return "", err
	}
	return accessToken.Token, nil
}

func (a *API) GetTagList(ctx context.Context, name string) ([]byte, error) {
	token, err := a.getAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	body, err := a.httpGet(ctx, fmt.Sprintf("%s?token=%s&name=%s", "api/v1/tags", token, name))
	if err != nil {
		return nil, err
	}
	return body, nil
}
