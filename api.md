# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#UserGetCurrentResponse">UserGetCurrentResponse</a>

Methods:

- <code title="get /v1/users/me">client.Users.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#UserService.GetCurrent">GetCurrent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#UserGetCurrentResponse">UserGetCurrentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lipsync

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncGenerate">LipsyncGenerate</a>

Methods:

- <code title="post /v1/lipsync/generate">client.Lipsync.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncGenerateParams">LipsyncGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncGenerate">LipsyncGenerate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/lipsync/generate-with-media">client.Lipsync.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncService.GenerateWithMedia">GenerateWithMedia</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncGenerateWithMediaParams">LipsyncGenerateWithMediaParams</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncGenerate">LipsyncGenerate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Requests

Response Types:

- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequest">LipsyncRequest</a>
- <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequestListResponse">LipsyncRequestListResponse</a>

Methods:

- <code title="get /v1/lipsync/requests/{id}">client.Lipsync.Requests.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequestService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequest">LipsyncRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/lipsync/requests">client.Lipsync.Requests.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequestService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequestListParams">LipsyncRequestListParams</a>) (\*<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2">chamelaion</a>.<a href="https://pkg.go.dev/github.com/chamelaion/chamelaion-go/v2#LipsyncRequestListResponse">LipsyncRequestListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Video Translate

Methods:

- `client.VideoTranslate.Generate(ctx, params)` starts translation from a remote video URL.
- `client.VideoTranslate.GenerateWithMedia(ctx, params)` uploads video and optional source/target SRT files.
- `client.VideoTranslate.Requests.Get(ctx, id)` retrieves one request.
- `client.VideoTranslate.Requests.List(ctx, params)` lists requests.
- `client.VideoTranslate.Voices.List(ctx, params)` lists available speaker voices.
