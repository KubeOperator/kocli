package auth

import (
	"encoding/json"
	"errors"
	"github.com/coreos/pkg/httputil"
	"github.com/go-resty/resty/v2"
	"kocli/pkg/config"
)

type Credential struct {
	Username string
	Password string
}

type SessionUser struct {
	Token string `json:"token"`
}

const authUrl = "/api/v1/auth/login"

func WithCredential(username, password string) error {
	c := config.GetConfig()
	client := resty.New()
	body, err := json.Marshal(Credential{
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	resp, err := client.SetHostURL(c.Server).
		R().
		SetHeader("Accept", httputil.JSONContentType).
		SetBody(body).
		Post(authUrl)
	if err != nil {
		return err
	}

	if resp.IsSuccess() {
		var user SessionUser
		err := json.Unmarshal(resp.Body(), &user)
		if err != nil {
			return err
		}
		err = config.SetToken(user.Token)
		if err != nil {
			return err
		}
	} else {
		return errors.New(string(resp.Body()))
	}
	return nil
}
