package csrf_test

import (
	"net/http"
	"time"

	"github.com/gogf/csrf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func ExampleNew() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(csrf.New())
		group.ALL("/csrf", func(r *ghttp.Request) {
			r.Response.Writeln(r.Method + ": " + r.RequestURI)
		})
	})
	s.SetPort(8199)
	s.Run()

	// Get http://localhost:8199/api.v2/csrf
	// get CSRF token in Cookie _csrf

	// Post http://localhost:8199/api.v2/csrf
	// invalid CSRF token in Header X-CSRF-Token
}

func ExampleNewWithCfg() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(csrf.NewWithCfg(csrf.Config{
			Cookie: &http.Cookie{
				Name:     "_csrf", // token name in cookie
				Secure:   true,
				SameSite: http.SameSiteNoneMode, // 自定义samesite
			},
			ExpireTime:      time.Hour * 24,
			TokenLength:     32,
			TokenRequestKey: "X-My-Token", // use this key to read token in request param
		}))
		group.ALL("/csrf", func(r *ghttp.Request) {
			r.Response.Writeln(r.Method + ": " + r.RequestURI)
		})
	})
	s.SetPort(8199)
	s.Run()

	// Get http://localhost:8199/api.v2/csrf
	// get CSRF token in Cookie _csrf

	// Post http://localhost:8199/api.v2/csrf
	// invalid CSRF token in Header X-My-Token
}
