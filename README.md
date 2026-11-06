# BrowserSolver Go SDK

The official Go SDK for the [BrowserSolver](https://www.browsersolver.com) web unlocker API.

Unlock, render, crawl and screenshot any website — anti-bot bypass and technology detection through one API.

The SDK exposes **9 typed methods**, one per documented endpoint, plus a generic `Request()`. It has no dependencies outside the standard library.

## Requirements

- Go 1.22 or newer

## Installation

```sh
go get github.com/browsersolver/browsersolver-go
```

## Get your API key

[Sign up for BrowserSolver](https://app.browsersolver.com/register) or sign in, then copy your API key from the [dashboard](https://app.browsersolver.com). `NewClient("", nil)` reads the `BROWSERSOLVER_API_KEY` environment variable.

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/browsersolver/browsersolver-go"
)

func main() {
	client, err := browsersolver.NewClient("", nil) // reads BROWSERSOLVER_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: "https://example.com"})
	if err != nil {
		log.Fatal(err)
	}
	data, err := response.Map()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
```

Parameter structs mirror the [API reference](https://app.browsersolver.com/openapi). Required fields are plain values; optional fields are pointers — use `browsersolver.String`, `browsersolver.Float`, `browsersolver.Int` and `browsersolver.Bool`.

Decode into your own types with `response.Decode(&value)`.

## Binary responses

Endpoints that return a file (image, PDF…) give you the bytes in `Body`; `IsJSON` tells both cases apart:

```go
response, err := client.WebsiteScreenshot(ctx, &browsersolver.WebsiteScreenshotParams{URL: "https://example.com", Format: browsersolver.String("png")})
if err != nil {
	log.Fatal(err)
}
if !response.IsJSON {
	err = os.WriteFile("screenshot.png", response.Body, 0o644)
}
```

## Async jobs

Any call accepts `is_async: true` (useful past the 60-second server limit): the response is then a `job_id`. `WaitForJob` polls `GET /v1/jobs/{id}` until a terminal status (`CompletedJob`, `Error`, `NotFound`, `BadRequest`):

```go
submitted, err := client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: "https://example.com", IsAsync: browsersolver.Bool(true)})
if err != nil {
	log.Fatal(err)
}
var accepted struct{ JobID string `json:"job_id"` }
if err := submitted.Decode(&accepted); err != nil {
	log.Fatal(err)
}
job, err := client.WaitForJob(ctx, accepted.JobID, &browsersolver.WaitForJobOptions{
	OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
})
```

## GET or POST

Methods use GET when the endpoint supports it (parameters in the query string); `WithMethod("POST")` sends them as a JSON body.

```go
response, err := client.WebsiteCrawler(ctx, &browsersolver.WebsiteCrawlerParams{Query: "https://books.toscrape.com/"}, browsersolver.WithMethod("POST"))
```

## Any endpoint

```go
response, err := client.Request(ctx, "GET", "/v1/usage", nil)
```

## Errors

```go
response, err := client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: "https://example.com"})
if apiError, ok := browsersolver.IsAPIError(err); ok {
	log.Printf("HTTP %d: %s", apiError.StatusCode, apiError.Message)
}
```

| Error | When |
| --- | --- |
| `*ConfigurationError` | Missing API key, invalid options or unsupported method. |
| `*SerializationError` | A required parameter is missing or a value can't be sent. |
| `*TransportError` | Network failure, timeout or cancelled context (`errors.Is(err, context.DeadlineExceeded)`). |
| `*APIError` | Non-2xx status (`StatusCode`, `Message`, `Response`). |
| `*ResponseDecodingError` | A success response could not be decoded. |

## Retries

Requests are retried up to 3 times after a network error, a `408`, `500`, `502`, `503` or `504`, or a rate limit — which the API returns as `401` (or `429`) with a `Rate limit exceeded` message. Per-key quota errors (`…: quota daily`) are not retried, nor are `400`, `402`, `404` or client-side timeouts. The delay starts at 1 second and grows 1.5× per retry (capped at 30 seconds); a `Retry-After` header wins.

```go
client, err := browsersolver.NewClient("", &browsersolver.ClientOptions{
	MaxRetries: 5,                      // negative disables retries
	RetryDelay: 500 * time.Millisecond,
})

// or for one request
response, err := client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: "https://example.com"}, browsersolver.WithMaxRetries(0))
```

## Batch requests

`browsersolver.Batch` runs a function per item with at most 10 calls in flight and returns one `BatchResult` per item, in input order. A failing item never stops the others.

```go
inputs := []string{"https://example.com", "https://www.browsersolver.com"}
results := browsersolver.Batch(ctx, inputs, 5, func(ctx context.Context, value string) (*browsersolver.Response, error) {
	return client.WebsiteWebunlocker(ctx, &browsersolver.WebsiteWebunlockerParams{Query: value})
})
for _, result := range results {
	if result.Err != nil {
		log.Printf("%s: %v", inputs[result.Index], result.Err)
		continue
	}
	fmt.Println(string(result.Value.Body))
}
```

| Constant | Default | Description |
| --- | --- | --- |
| `DefaultConcurrency` | `10` | Tasks run at the same time by batch |
| `DefaultTimeout` | `120 * time.Second` | Request timeout |
| `MaxRetries` | `3` | Retry attempts after a network error, a 408/5xx or a rate limit |
| `RetryBackoffFactor` | `1.5` | Exponential backoff multiplier |
| `DefaultRetryDelay` | `time.Second` | Delay before the first retry |

## Timeouts and custom HTTP clients

```go
client, err := browsersolver.NewClient("", &browsersolver.ClientOptions{
	Timeout:    60 * time.Second, // default 120 s; negative disables
	HTTPClient: &http.Client{Transport: myTransport},
})
```

Every method also honours the deadline and cancellation of the `context.Context` you pass.

## Endpoints

9 methods, generated from the [OpenAPI document](https://app.browsersolver.com/openapi).

| Method | HTTP | Description |
| --- | --- | --- |
| `Jobs` | GET `/v1/jobs/{id}` | Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: Completed… |
| `Stats` | GET `/v1/stats` | Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 mea… |
| `Usage` | GET `/v1/usage` | Retrieve real-time usage statistics and credit balance for your BrowserSolver account. |
| `WebsiteAntibot` | GET `/v1/website/antibot` | Single-pass scan that names the bot-protection vendor in front of a URL so you can choose the right BrowserSolver fetch tier next. |
| `WebsiteCrawler` | GET / POST `/v1/website/crawler` | Fast HTTP pull of public HTML with no browser session. Ideal when the first response already contains the content you need. |
| `WebsiteRendering` | GET `/v1/website/rendering` | Run a managed browser session against a public URL and collect the HTML after client-side scripts finish. Prefer Crawler for static markup;… |
| `WebsiteScreenshot` | GET / POST `/v1/website/screenshot` | Turn any public URL into a PNG, JPEG, WebP, or PDF asset rendered by BrowserSolver. No local Chromium fleet to operate. |
| `WebsiteTechnology` | GET / POST `/v1/website/technology` | Map CMS, frameworks, analytics, CDN, and security tooling exposed by a public URL. Choose simple lists or expert metadata with confidence s… |
| `WebsiteWebunlocker` | GET / POST `/v1/website/webunlocker` | BrowserSolver's highest HTML fetch tier for retail and marketplace hosts on your account whitelist. Solves persistent bot challenges that s… |

## All BrowserSolver SDKs

| SDK | Package | Install |
| --- | --- | --- |
| Node.js | [`@browsersolver/sdk`](https://www.npmjs.com/package/@browsersolver/sdk) | `npm install @browsersolver/sdk` |
| Python | [`browsersolver`](https://pypi.org/project/browsersolver/) | `pip install browsersolver` |
| Go | [`browsersolver-go`](https://pkg.go.dev/github.com/browsersolver/browsersolver-go) | `go get github.com/browsersolver/browsersolver-go` |
| .NET | [`BrowserSolver`](https://www.nuget.org/packages/BrowserSolver) | `dotnet add package BrowserSolver` |
| Ruby | [`browsersolver`](https://rubygems.org/gems/browsersolver) | `gem install browsersolver` |
| PHP | [`browsersolver/sdk`](https://packagist.org/packages/browsersolver/sdk) | `composer require browsersolver/sdk` |
| Java | [`browsersolver-java`](https://github.com/browsersolver/browsersolver-java) | build from source, see its README |
| CLI | [`@browsersolver/cli`](https://www.npmjs.com/package/@browsersolver/cli) | `npm install -g @browsersolver/cli` |
| MCP server | [`@browsersolver/mcp`](https://www.npmjs.com/package/@browsersolver/mcp) | `npx -y @browsersolver/mcp` |

## License

MIT — see [LICENSE](LICENSE).
