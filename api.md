# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#UserGetCurrentResponse">UserGetCurrentResponse</a>

Methods:

- <code title="get /v1/users/me">client.Users.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#UserService.GetCurrent">GetCurrent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#UserGetCurrentResponse">UserGetCurrentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lipsync

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncGenerate">LipsyncGenerate</a>

Methods:

- <code title="post /v1/lipsync/generate">client.Lipsync.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncGenerateParams">LipsyncGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncGenerate">LipsyncGenerate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/lipsync/generate-with-media">client.Lipsync.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncService.GenerateWithMedia">GenerateWithMedia</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncGenerateWithMediaParams">LipsyncGenerateWithMediaParams</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncGenerate">LipsyncGenerate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Requests

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequest">LipsyncRequest</a>
- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequestListResponse">LipsyncRequestListResponse</a>

Methods:

- <code title="get /v1/lipsync/requests/{id}">client.Lipsync.Requests.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequest">LipsyncRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/lipsync/requests">client.Lipsync.Requests.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequestListParams">LipsyncRequestListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go#LipsyncRequestListResponse">LipsyncRequestListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
