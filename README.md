# gf-csrf

## How to use

```go
package main

import (
	"net/http"
	"time"

	"github.com/gogf/gf-csrf/csrf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// default cfg
func main() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(csrf.NewCSRF())
		group.ALL("/csrf", func(r *ghttp.Request) {
			r.Response.Writeln(r.Method + ": " + r.RequestURI)
		})
	})
	s.SetPort(8199)
	s.Run()
}
```

```go
package main

import (
	"net/http"
	"time"

	"github.com/gogf/gf-csrf/csrf"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// set cfg
func main() {
	s := g.Server()
	s.Group("/api.v2", func(group *ghttp.RouterGroup) {
		group.Middleware(csrf.NewCSRFWithCfg(csrf.CSRFConfig{
			Cookie: &http.Cookie{
				Name: "_csrf",// token name in cookie
			},
			ExpireTime:      time.Hour * 24,
			TokenLength:     32,
			TokenRequestKey: "X-CSRF-Token",// use this key to read token in request param
		}))
		group.ALL("/csrf", func(r *ghttp.Request) {
			r.Response.Writeln(r.Method + ": " + r.RequestURI)
		})
	})
	s.SetPort(8199)
	s.Run()
}
```



## Check effect by request(GET and POST)

http://localhost:8199/api.v2/csrf

You can set the token in request with param(Router < Query < Body < Form < Custom < Header)