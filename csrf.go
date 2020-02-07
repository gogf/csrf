package csrf

import (
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/grand"
)

// CSRFConfig CSRFConfig struct
type CSRFConfig struct {
	TokenLength     int
	TokenRequestKey string
	ExpireTime      time.Duration
	Cookie          *http.Cookie
}

var (
	// DefaultCSRFConfig is the default CSRF middleware config.
	DefaultCSRFConfig = CSRFConfig{
		Cookie: &http.Cookie{
			Name: "_csrf",
		},
		ExpireTime:      time.Hour * 24,
		TokenLength:     32,
		TokenRequestKey: "X-CSRF-Token",
	}
)

// NewCSRF Create CSRF middleware (with default configuration)
//
// createTime: 2020年01月21日 17:03:26
//
// author: hailaz
func NewCSRF() func(r *ghttp.Request) {
	return NewCSRFWithCfg(DefaultCSRFConfig)
}

// NewCSRFWithCfg Create CSRF middleware (with incoming configuration)
//
// createTime: 2020年01月20日 17:51:06
//
// author: hailaz
func NewCSRFWithCfg(cfg CSRFConfig) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {

		// Read the token in the request cookie
		tokenInCookie := r.Cookie.Get(cfg.Cookie.Name)
		if tokenInCookie == "" {
			// Generate a random token
			tokenInCookie = grand.Str(cfg.TokenLength)
		}

		// Read the token attached to the request
		// Read priority: Router < Query < Body < Form < Custom < Header
		tokenInRequestData := r.Header.Get(cfg.TokenRequestKey)
		if tokenInRequestData == "" {
			tokenInRequestData = r.GetString(cfg.TokenRequestKey)
		}

		switch r.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
			// No verification required
		default:
			// Authentication token
			if !strings.EqualFold(tokenInCookie, tokenInRequestData) {
				r.Response.WriteStatusExit(http.StatusForbidden, "invalid csrf token")
				return
			}
		}

		// Set cookie timeout
		cfg.Cookie.Expires = time.Now().Add(cfg.ExpireTime)
		cfg.Cookie.Value = tokenInCookie

		// Set cookies in response
		http.SetCookie(r.Response.RawWriter(), cfg.Cookie)
		r.Middleware.Next()
	}
}