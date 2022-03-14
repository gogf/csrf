# csrf

csrf middleware for goframe web module

## How to use

```go
// v2 for gf/v2
go get github.com/gogf/csrf/v2

// v1 for gf
go get github.com/gogf/csrf
```

[example](example_test.go)

## Check the effect by web request(GET before POST)

http://127.0.0.1:8199/api.v2/csrf

You can get token value(default name _csrf) in cookies

You can set the token in request with param(Router < Query < Body < Form < Custom < Header)
