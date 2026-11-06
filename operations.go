// Code generated from the BrowserSolver OpenAPI document. DO NOT EDIT.

package browsersolver

import "context"

// Jobs — Jobs. Poll the status, progress (0–100) and result of a job started with is_async=true. Polling is free. Poll until status is terminal: CompletedJob, Error, NotFound or BadRequest.
//
// GET /v1/jobs/{id}
func (c *Client) Jobs(ctx context.Context, params *JobsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["jobs"], params.params(), options)
}

// Stats — Stats. Infrastructure health for an engine or endpoint over the last hour. Not scoped to your account or tenant. Enabled per account: HTTP 403 means contact support to activate it.
//
// GET /v1/stats
func (c *Client) Stats(ctx context.Context, params *StatsParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["stats"], params.params(), options)
}

// Usage — Usage. Retrieve real-time usage statistics and credit balance for your BrowserSolver account.
//
// GET /v1/usage
func (c *Client) Usage(ctx context.Context, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["usage"], nil, options)
}

// WebsiteAntibot — Website Antibot. Single-pass scan that names the bot-protection vendor in front of a URL so you can choose the right BrowserSolver fetch tier next. Credit cost: 1 credit per call.
//
// GET /v1/website/antibot
func (c *Client) WebsiteAntibot(ctx context.Context, params *WebsiteAntibotParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_antibot"], params.params(), options)
}

// WebsiteCrawler — Website Crawler. Fast HTTP pull of public HTML with no browser session. Ideal when the first response already contains the content you need. Credit cost: 1 credit per call.
//
// GET / POST /v1/website/crawler
func (c *Client) WebsiteCrawler(ctx context.Context, params *WebsiteCrawlerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_crawler"], params.params(), options)
}

// WebsiteRendering — Website Rendering. Run a managed browser session against a public URL and collect the HTML after client-side scripts finish. Prefer Crawler for static markup; escalate to WebUnlocker on approved hosts when challenges persist. Credit cost: 2 credits per call.
//
// GET /v1/website/rendering
func (c *Client) WebsiteRendering(ctx context.Context, params *WebsiteRenderingParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_rendering"], params.params(), options)
}

// WebsiteScreenshot — Website Screenshot. Turn any public URL into a PNG, JPEG, WebP, or PDF asset rendered by BrowserSolver. No local Chromium fleet to operate. Credit cost: 2 credits per call.
//
// GET / POST /v1/website/screenshot
func (c *Client) WebsiteScreenshot(ctx context.Context, params *WebsiteScreenshotParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_screenshot"], params.params(), options)
}

// WebsiteTechnology — Website Technology. Map CMS, frameworks, analytics, CDN, and security tooling exposed by a public URL. Choose simple lists or expert metadata with confidence scores. Credit cost: 1 credit per call.
//
// GET / POST /v1/website/technology
func (c *Client) WebsiteTechnology(ctx context.Context, params *WebsiteTechnologyParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_technology"], params.params(), options)
}

// WebsiteWebunlocker — Website WebUnlocker. BrowserSolver's highest HTML fetch tier for retail and marketplace hosts on your account whitelist. Solves persistent bot challenges that still block Rendering. Credit cost: 3 credits per call.
//
// GET / POST /v1/website/webunlocker
func (c *Client) WebsiteWebunlocker(ctx context.Context, params *WebsiteWebunlockerParams, options ...RequestOption) (*Response, error) {
	return c.call(ctx, endpoints["website_webunlocker"], params.params(), options)
}

// endpoints lists every documented operation by its snake_case name.
var endpoints = map[string]endpoint{
	"jobs":                {path: "/v1/jobs/{id}", methods: []string{"GET"}, json: true, required: []string{"id"}},
	"stats":               {path: "/v1/stats", methods: []string{"GET"}, json: true, required: []string{}},
	"usage":               {path: "/v1/usage", methods: []string{"GET"}, json: true, required: []string{}},
	"website_antibot":     {path: "/v1/website/antibot", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"website_crawler":     {path: "/v1/website/crawler", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_rendering":   {path: "/v1/website/rendering", methods: []string{"GET"}, json: true, required: []string{"query"}},
	"website_screenshot":  {path: "/v1/website/screenshot", methods: []string{"GET", "POST"}, json: false, required: []string{"url"}},
	"website_technology":  {path: "/v1/website/technology", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
	"website_webunlocker": {path: "/v1/website/webunlocker", methods: []string{"GET", "POST"}, json: true, required: []string{"query"}},
}
