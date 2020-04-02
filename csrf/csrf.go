package csrf

import (
	"net/http"
	"strings"
	"time"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/grand"
)

// Config is the configuration struct for CSRF feature.
type Config struct {
	TokenLength     int
	TokenRequestKey string
	ExpireTime      time.Duration
	Cookie          *http.Cookie
}

var (
	// DefaultCSRFConfig is the default CSRF middleware config.
	DefaultCSRFConfig = Config{
		Cookie: &http.Cookie{
			Name: "_csrf",
		},
		ExpireTime:      time.Hour * 24,
		TokenLength:     32,
		TokenRequestKey: "X-CSRF-Token",
	}
)

// New creates and returns a CSRF middleware with default configuration.
func New() func(r *ghttp.Request) {
	return NewWithCfg(DefaultCSRFConfig)
}

// NewWithCfg creates and returns a CSRF middleware with incoming configuration.
func NewWithCfg(cfg Config) func(r *ghttp.Request) {
	return func(r *ghttp.Request) {

		// Read the token in the request cookie
		tokenInCookie := r.Cookie.Get(cfg.Cookie.Name)
		if tokenInCookie == "" {
			// Generate a random token
			tokenInCookie = grand.S(cfg.TokenLength)
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
				r.Response.WriteStatusExit(http.StatusForbidden, "invalid CSRF token")
				return
			}
		}

		// Set cookie timeout
		cfg.Cookie.Expires = time.Now().Add(cfg.ExpireTime)
		cfg.Cookie.Value = tokenInCookie

		// Set cookie in response
		http.SetCookie(r.Response.RawWriter(), cfg.Cookie)
		r.Middleware.Next()
	}
}
