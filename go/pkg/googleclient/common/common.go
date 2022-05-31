package common

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/tasks/v1"
)

func CreateGoogleClient(file string) *http.Client {
	return http.DefaultClient
}

func getGoogleOauth2ConfigFromFile(file string) (*oauth2.Config, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.WithMessage(err, "read google oauth2 config")
	}

	config, err := google.ConfigFromJSON(b, tasks.TasksScope)
	if err != nil {
		return nil, errors.WithMessage(err, "get google oauth2 config from json")
	}

	return config, nil
}

func getTokenFromWeb(cfg *oauth2.Config) (*oauth2.Token, error) {
	cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	token, err := cfg.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func createClientFromOauth2Config(cfg *oauth2.Config) (*http.Client, error) {
	token, err := getTokenFromWeb(cfg)
	if err != nil {
		return nil, err
	}

	return cfg.Client(context.Background(), token), nil
}
