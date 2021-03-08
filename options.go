package kuaishou

import (
	"net/http"
	"time"
)

// Options options
type Options struct {
	APPID       string
	AppSecret   string
	CallBackURL string
}

type httpSettings struct {
	headers map[string]string
	cookies []*http.Cookie
	close   bool
	timeout time.Duration
}

// HTTPOption http request配置
type HTTPOption func(s *httpSettings)
