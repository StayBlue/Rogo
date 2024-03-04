package Rogo

import (
	"context"
	"github.com/PuerkitoBio/rehttp"
	"github.com/carlmjohnson/requests"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	config func(*requests.Builder)
}

var ctx = context.Background()

func tokenTransport(token *string) requests.Transport {
	return requests.RoundTripFunc(func(req *http.Request) (res *http.Response, err error) {
		req.Header.Set("X-CSRF-TOKEN", *token)
		return http.DefaultTransport.RoundTrip(req)
	})
}

func tokenHandler(token *string) rehttp.RetryFn {
	return func(attempt rehttp.Attempt) bool {
		resp := attempt.Response
		_token := resp.Header.Get("X-CSRF-TOKEN")
		if resp.StatusCode == http.StatusForbidden && _token != "" {
			*token = _token
			return true
		}
		return false
	}
}

func (c *Client) getRequest(endpoint string) *requests.Builder {
	return requests.New(c.config).
		Hostf("%s.roblox.com", endpoint)
}

func NewClient(token string) *Client {
	jar := requests.NewCookieJar()

	client := &http.Client{
		Jar: jar,
	}

	var _token string
	config := func(rb *requests.Builder) {
		rb.
			Cookie(".ROBLOSECURITY", token).
			UserAgent("Rogo").
			Transport(rehttp.NewTransport(tokenTransport(&_token), rehttp.RetryAll(tokenHandler(&_token), rehttp.RetryMaxRetries(5)), rehttp.ConstDelay(50*time.Millisecond)))
	}

	return &Client{
		client: client,
		config: config,
	}
}
