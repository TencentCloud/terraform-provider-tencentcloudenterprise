package connectivity

import (
	"net/http"
	"os"
)

const REQUEST_CLIENT = "TENCENTCLOUD_API_REQUEST_CLIENT"

var ReqClient = "Terraform-latest"

func SetReqClient(name string) {
	if name == "" {
		return
	}
	ReqClient = name
}

type LogRoundTripper struct {
}

func (me *LogRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	client := os.Getenv(REQUEST_CLIENT)
	if client == "" {
		client = ReqClient
	}
	request.Header.Set("X-TC-RequestClient", client)
	return http.DefaultTransport.RoundTrip(request)
}
