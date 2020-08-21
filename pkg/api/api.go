package api

import (
	"fmt"
	"github.com/coreos/pkg/httputil"
	"gopkg.in/resty.v1"
)

const server = "localhost:8080"

var token = "123456"

func GetApiClient(path string) *resty.Client {
	client := resty.New()
	return client.SetAuthToken(token).
		SetHeader("Accept", httputil.JSONContentType).
		SetHostURL(fmt.Sprintf("http://%s/api/v1/%s", server, path))
}

