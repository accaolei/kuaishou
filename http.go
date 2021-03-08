package kuaishou

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

// HTTPClient is the interface that do http request
type HTTPClient interface {
	// Get http的get请求
	Get(ctx context.Context, url string, options ...HTTPOption) ([]byte, error)
	// Post post请求
	Post(ctx context.Context, url string, body []byte, options ...HTTPOption) ([]byte, error)
}

type apiClient struct {
	client  *http.Client
	timeout time.Duration
}

func (c *apiClient) do(ctx context.Context, req *http.Request, options ...HTTPOption) ([]byte, error) {
	settings := &httpSettings{timeout: c.timeout}
	if len(options) != 0 {
		settings.headers = make(map[string]string)
		for _, f := range options {
			f(settings)
		}
	}
	if len(settings.headers) != 0 {
		for _, v := range settings.cookies {
			req.AddCookie(v)
		}
	}
	if settings.close {
		req.Close = true
	}
	ctx, cancel := context.WithTimeout(ctx, settings.timeout)
	defer cancel()
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		io.Copy(ioutil.Discard, resp.Body)
		return nil, fmt.Errorf("error http code:%d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *apiClient) Get(ctx context.Context, url string, options ...HTTPOption) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, options...)
}

func (c *apiClient) Post(ctx context.Context, url string, body []byte, options ...HTTPOption) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.do(ctx, req, options...)
}

// NewHTTPClient new http client
func NewHTTPClient(tlsCfg ...*tls.Config) HTTPClient {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
	}
	return &apiClient{
		client: &http.Client{
			Transport: t,
		},
		timeout: defaultTimeout,
	}
}
