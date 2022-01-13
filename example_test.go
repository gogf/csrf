package csrf_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gogf/csrf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func ExampleNew() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(csrf.New())
		group.ALL("/csrf", func(r *ghttp.Request) {
			reqInfo := "Request Info:\n"
			reqInfo += fmt.Sprintf("%s\n", r.Method+": "+r.RequestURI)
			reqInfo += fmt.Sprintf("Cookies: %v\n", r.Cookies())
			reqInfo += fmt.Sprintf("Header: %v\n", r.Header)
			reqInfo += fmt.Sprintf("Query: %v\n", r.URL.Query())
			glog.Debug(r.Context(), reqInfo)
			r.Response.Writeln(reqInfo)
		})
	})
	s.SetAddr("127.0.0.1:8199")
	s.Run()

	// Get http://127.0.0.1:8199/api.v2/csrf
	// get CSRF token in Cookie _csrf

	// Post http://127.0.0.1:8199/api.v2/csrf
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
			reqInfo := "Request Info:\n"
			reqInfo += fmt.Sprintf("%s\n", r.Method+": "+r.RequestURI)
			reqInfo += fmt.Sprintf("Cookies: %v\n", r.Cookies())
			reqInfo += fmt.Sprintf("Header: %v\n", r.Header)
			reqInfo += fmt.Sprintf("Query: %v\n", r.URL.Query())
			glog.Debug(r.Context(), reqInfo)
			r.Response.Writeln(reqInfo)
		})
	})
	s.SetAddr("127.0.0.1:8199")
	s.Run()
	// Get http://127.0.0.1:8199/api.v2/csrf
	// get CSRF token in Cookie _csrf

	// Post http://127.0.0.1:8199/api.v2/csrf
	// invalid CSRF token in Header X-My-Token
}
