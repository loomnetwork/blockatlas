package loom

import (
	"fmt"
	"net/http"
	"strconv"

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
	fmt.Printf("\nLOOM CURL : %+v\n", c)
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

//TODO: need to implement this endpoint to loomchain
func (c *Client) GetPool() (result StakingPool, err error) {
	return result, c.Request.Get(&result, c.URL, "query/staking/pool", nil)
}

func (c *Client) GetRate() (float64, error) {
	var result string

	err := c.Request.Get(&result, c.URL, "query/staking/rate", nil)
	if err != nil {
		return 0, err
	}

	s, err := strconv.ParseFloat(result, 32)

	return s, err
}

func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var block Block
	err = c.Request.Get(&block, c.URL, "query/getblockheight", nil)
	if err != nil {
		return num, err
	}
	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)

	if err != nil {
		return num, err
	}

	return num, nil
}
