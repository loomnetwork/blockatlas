package loom

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
)

// Client - the HTTP client
type Client struct {
	Request blockatlas.Request
	URL     string
	RpcURL  string
}

func InitClient(URL, RpcURL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		URL:    URL,
		RpcURL: RpcURL,
	}
}

func (c *Client) GetValidators() (validators []Validator, err error) {
	var info = struct {
		JSONRPC string `json:"jsonrpc"`
		ID      string `json:"id"`
		Result  struct {
			Validators []Validator `json:"validators,omitempty"`
		} `json:"result"`
	}{}
	err = c.Request.Get(&info, c.URL, "query/getvalidators", nil)
	if err != nil {
		logrus.WithError(err).Errorf("LOOM : Failed to get validators")
		return info.Result.Validators, err
	}
	return info.Result.Validators, err
}
