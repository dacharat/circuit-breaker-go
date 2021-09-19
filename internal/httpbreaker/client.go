package httpbreaker

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sony/gobreaker"
)

func Init() {

}

type Client struct {
	client http.Client
	cb     *gobreaker.CircuitBreaker
}

func NewDefaultClient() Client {
	return Client{
		client: http.Client{},
		cb:     NewCircuitBreaker(),
	}
}

func (c *Client) Get(ctx context.Context, url string, header http.Header) ([]byte, error) {

	body, err := c.cb.Execute(func() (interface{}, error) {
		resp, err := c.get(ctx, url, header)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= http.StatusInternalServerError {
			return body, errors.New("internal server error")
		}

		return body, nil
	})

	if err != nil {
		return nil, err
	}

	return body.([]byte), nil

	// return c.get(ctx, url, header)
}

func (c *Client) get(ctx context.Context, url string, header http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// header = http.Header{}
	// header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header = header

	return c.client.Do(req)
}

func (c *Client) Post(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	// header = http.Header{}
	// header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header = header

	return c.client.Do(req)
}

func (c *Client) Put(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	// header = http.Header{}
	// header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header = header

	return c.client.Do(req)
}

func (c *Client) Delete(ctx context.Context, url string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}

	// header = http.Header{}
	// header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header = header

	return c.client.Do(req)
}
