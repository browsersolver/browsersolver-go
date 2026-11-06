// Code generated from the BrowserSolver OpenAPI document. DO NOT EDIT.

package browsersolver

// JobsParams holds the parameters of Jobs (GET /v1/jobs/{id}).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type JobsParams struct {
	// The job_id returned by an endpoint called with is_async=true.
	ID string `json:"id" api:"id,required"`
}

func (p *JobsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 1)
	if p.ID != "" {
		values = append(values, param{"id", p.ID})
	}
	return values
}

// StatsParams holds the parameters of Stats (GET /v1/stats).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type StatsParams struct {
	// API family to check, for example website. Provide at least engine or path.
	Engine *string `json:"engine" api:"engine"`
	// Specific endpoint to check, for example /v1/website/screenshot. If you also pass engine, use an endpoint from that API.
	Path *string `json:"path" api:"path"`
}

func (p *StatsParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 2)
	if p.Engine != nil {
		values = append(values, param{"engine", *p.Engine})
	}
	if p.Path != nil {
		values = append(values, param{"path", *p.Path})
	}
	return values
}

// WebsiteAntibotParams holds the parameters of WebsiteAntibot (GET /v1/website/antibot).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type WebsiteAntibotParams struct {
	// Full URL or bare domain to scan for bot-protection vendors.
	Query string `json:"query" api:"query,required"`
	// Upper bound for waiting on the target HTTP response, expressed in seconds (defaults to 15).
	Timeout *float64 `json:"timeout" api:"timeout"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteAntibotParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Timeout != nil {
		values = append(values, param{"timeout", *p.Timeout})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteCrawlerParams holds the parameters of WebsiteCrawler (GET / POST /v1/website/crawler).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type WebsiteCrawlerParams struct {
	// Public URL to retrieve with the Crawler tier (plain HTTP, no JavaScript). Include http:// or https://.
	Query string `json:"query" api:"query,required"`
	// Whether BrowserSolver should chase HTTP redirects for this fetch. Default: true.
	AllowRedirects *bool `json:"allow_redirects" api:"allow_redirects"`
	// Choose between processed markup and the exact bytes from the origin. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteCrawlerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.AllowRedirects != nil {
		values = append(values, param{"allow_redirects", *p.AllowRedirects})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteRenderingParams holds the parameters of WebsiteRendering (GET /v1/website/rendering).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type WebsiteRenderingParams struct {
	// Public URL to load inside BrowserSolver's headless browser. Include http:// or https://.
	Query string `json:"query" api:"query,required"`
	// Extra pause, in seconds, after the page load event before BrowserSolver returns HTML. Helpful when widgets animate or hydrate late. Default: 0.
	WaitInSeconds *float64 `json:"wait_in_seconds" api:"wait_in_seconds"`
	// CSS or XPath selector that must appear in the DOM before HTML is returned. The managed browser waits until the node exists (or until timeout).
	WaitFor *string `json:"wait_for" api:"wait_for"`
	// Skip ad-network requests while the page renders. Default: false.
	BlockAds *bool `json:"block_ads" api:"block_ads"`
	// Hard ceiling, in seconds, for the render session. When reached, BrowserSolver returns whatever HTML is available. Default: 30.
	Timeout *float64 `json:"timeout" api:"timeout"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteRenderingParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 8)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.WaitInSeconds != nil {
		values = append(values, param{"wait_in_seconds", *p.WaitInSeconds})
	}
	if p.WaitFor != nil {
		values = append(values, param{"wait_for", *p.WaitFor})
	}
	if p.BlockAds != nil {
		values = append(values, param{"block_ads", *p.BlockAds})
	}
	if p.Timeout != nil {
		values = append(values, param{"timeout", *p.Timeout})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteScreenshotParams holds the parameters of WebsiteScreenshot (GET / POST /v1/website/screenshot).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type WebsiteScreenshotParams struct {
	// Page address to render into an image or PDF. Must include a protocol.
	URL string `json:"url" api:"url,required"`
	// File type returned in the response body. Allowed values: jpeg, jpg, pdf, png, webp. Default: "png".
	Format *string `json:"format" api:"format"`
	// Swap the desktop viewport for a preset handset or tablet fingerprint. Allowed values: galaxy_fold4, galaxy_s23, galaxy_s23_ultra, galaxy_s8, huawei_mate_50_pro, huawei_p60_pro, ipad, iphone_12, iphone_13_mini, iphone_13_pro_max, iphone_14_pro, iphone_14_pro_max, iphone_x, pixel_5, pixel_6a, pixel_7…
	Device *string `json:"device" api:"device"`
	// Extend the capture beyond the viewport so the entire scrollable document is included. Default: false.
	FullPage *bool `json:"full_page" api:"full_page"`
	// Drive a scroll pass first so deferred content mounts before a tall capture. Default: false.
	FullPageScroll *bool `json:"full_page_scroll" api:"full_page_scroll"`
	// Milliseconds spent scrolling prior to a full-height capture. Default: 400.
	FullPageScrollDuration *float64 `json:"full_page_scroll_duration" api:"full_page_scroll_duration"`
	// Width of the synthetic desktop window, in pixels. Default: 1280.
	ViewportWidth *float64 `json:"viewport_width" api:"viewport_width"`
	// Height of the synthetic desktop window, in pixels. Default: 1024.
	ViewportHeight *float64 `json:"viewport_height" api:"viewport_height"`
	// Pixel density multiplier. Raise it for retina-quality output; files grow accordingly. Default: 1.
	ScaleFactor *float64 `json:"scale_factor" api:"scale_factor"`
	// Lossy compression level for soft formats. Default: "80".
	ImageQuality *string `json:"image_quality" api:"image_quality"`
	// Limit the shot to one DOM node matched by this CSS selector.
	Selector *string `json:"selector" api:"selector"`
	// Serve a stored capture on repeat requests when caching is enabled. Default: false.
	Cache *bool `json:"cache" api:"cache"`
	// Lifetime of cached captures in seconds. Default: 2592000.
	CacheTTL *float64 `json:"cache_ttl" api:"cache_ttl"`
	// Hide matching CSS selectors (modals, promo bars, etc.) before shooting.
	RemoveSelectors *string `json:"remove_selectors" api:"remove_selectors"`
	// Best-effort removal of cookie consent layers before capture. Default: false.
	RemoveCookieBanners *bool `json:"remove_cookie_banners" api:"remove_cookie_banners"`
	// Best-effort stripping of ad slots before capture. Default: false.
	RemoveAds *bool `json:"remove_ads" api:"remove_ads"`
	// Resource kinds to discard during load, as a comma list (example: image,stylesheet,font). Allowed values: document, eventsource, fetch, font, image, manifest, media, other, script, stylesheet, texttrack, websocket, xhr.
	BlockResources *string `json:"block_resources" api:"block_resources"`
	// Substring patterns that should never be fetched (example: analytics,tracking,advertisement). Hosts, paths, or fragments such as .example.com/ work.
	BlockUrls *string `json:"block_urls" api:"block_urls"`
	// Additional settle time in seconds after the readiness event. Default: 0.
	Delay *float64 `json:"delay" api:"delay"`
	// Which browser lifecycle event marks the page ready. Allowed values: domcontentloaded, load, networkidle0, networkidle2.
	WaitUntil []string `json:"wait_until" api:"wait_until"`
	// Hold capture until this CSS node is present.
	WaitForSelector *string `json:"wait_for_selector" api:"wait_for_selector"`
	// Route the browser through an intermediate proxy.
	Proxy *string `json:"proxy" api:"proxy"`
	// Secret key used to authenticate uploads.
	S3SecretKey *string `json:"s3_secret_key" api:"s3_secret_key"`
	// Access key used to authenticate uploads.
	S3AccessKeyID *string `json:"s3_access_key_id" api:"s3_access_key_id"`
	// Destination bucket for the binary asset.
	S3Bucket *string `json:"s3_bucket" api:"s3_bucket"`
	// Key under which the object is stored (generated automatically if blank).
	S3ObjectKey *string `json:"s3_object_key" api:"s3_object_key"`
	// Cloud region of the bucket.
	S3Region *string `json:"s3_region" api:"s3_region"`
	// Override the default AWS endpoint for S3-compatible backends. Leave blank for AWS. R2 example: https://<accountId>.r2.cloudflarestorage.com.
	StorageEndpoint *string `json:"storage_endpoint" api:"storage_endpoint"`
	// Upload the binary and include the public object URL in the API response. Default: false.
	S3URL *bool `json:"s3_url" api:"s3_url"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteScreenshotParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 32)
	if p.URL != "" {
		values = append(values, param{"url", p.URL})
	}
	if p.Format != nil {
		values = append(values, param{"format", *p.Format})
	}
	if p.Device != nil {
		values = append(values, param{"device", *p.Device})
	}
	if p.FullPage != nil {
		values = append(values, param{"full_page", *p.FullPage})
	}
	if p.FullPageScroll != nil {
		values = append(values, param{"full_page_scroll", *p.FullPageScroll})
	}
	if p.FullPageScrollDuration != nil {
		values = append(values, param{"full_page_scroll_duration", *p.FullPageScrollDuration})
	}
	if p.ViewportWidth != nil {
		values = append(values, param{"viewport_width", *p.ViewportWidth})
	}
	if p.ViewportHeight != nil {
		values = append(values, param{"viewport_height", *p.ViewportHeight})
	}
	if p.ScaleFactor != nil {
		values = append(values, param{"scale_factor", *p.ScaleFactor})
	}
	if p.ImageQuality != nil {
		values = append(values, param{"image_quality", *p.ImageQuality})
	}
	if p.Selector != nil {
		values = append(values, param{"selector", *p.Selector})
	}
	if p.Cache != nil {
		values = append(values, param{"cache", *p.Cache})
	}
	if p.CacheTTL != nil {
		values = append(values, param{"cache_ttl", *p.CacheTTL})
	}
	if p.RemoveSelectors != nil {
		values = append(values, param{"remove_selectors", *p.RemoveSelectors})
	}
	if p.RemoveCookieBanners != nil {
		values = append(values, param{"remove_cookie_banners", *p.RemoveCookieBanners})
	}
	if p.RemoveAds != nil {
		values = append(values, param{"remove_ads", *p.RemoveAds})
	}
	if p.BlockResources != nil {
		values = append(values, param{"block_resources", *p.BlockResources})
	}
	if p.BlockUrls != nil {
		values = append(values, param{"block_urls", *p.BlockUrls})
	}
	if p.Delay != nil {
		values = append(values, param{"delay", *p.Delay})
	}
	if p.WaitUntil != nil {
		values = append(values, param{"wait_until", p.WaitUntil})
	}
	if p.WaitForSelector != nil {
		values = append(values, param{"wait_for_selector", *p.WaitForSelector})
	}
	if p.Proxy != nil {
		values = append(values, param{"proxy", *p.Proxy})
	}
	if p.S3SecretKey != nil {
		values = append(values, param{"s3_secret_key", *p.S3SecretKey})
	}
	if p.S3AccessKeyID != nil {
		values = append(values, param{"s3_access_key_id", *p.S3AccessKeyID})
	}
	if p.S3Bucket != nil {
		values = append(values, param{"s3_bucket", *p.S3Bucket})
	}
	if p.S3ObjectKey != nil {
		values = append(values, param{"s3_object_key", *p.S3ObjectKey})
	}
	if p.S3Region != nil {
		values = append(values, param{"s3_region", *p.S3Region})
	}
	if p.StorageEndpoint != nil {
		values = append(values, param{"storage_endpoint", *p.StorageEndpoint})
	}
	if p.S3URL != nil {
		values = append(values, param{"s3_url", *p.S3URL})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteTechnologyParams holds the parameters of WebsiteTechnology (GET / POST /v1/website/technology).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type WebsiteTechnologyParams struct {
	// Public URL to fingerprint for CMS, frameworks, analytics, CDN, and security tooling. Include http:// or https://.
	Query string `json:"query" api:"query,required"`
	// How much detail to include in the stack report. Default: "expert".
	Mode *string `json:"mode" api:"mode"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteTechnologyParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 5)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.Mode != nil {
		values = append(values, param{"mode", *p.Mode})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}

// WebsiteWebunlockerParams holds the parameters of WebsiteWebunlocker (GET / POST /v1/website/webunlocker).
// Required fields are plain values; optional fields are pointers (use browsersolver.String, browsersolver.Float, browsersolver.Bool) or nil slices and maps.
type WebsiteWebunlockerParams struct {
	// Public URL to fetch with the WebUnlocker tier. Include http:// or https://. The hostname must be on your BrowserSolver whitelist.
	Query string `json:"query" api:"query,required"`
	// Whether the unlock session should chase HTTP redirects. Default: true.
	AllowRedirects *bool `json:"allow_redirects" api:"allow_redirects"`
	// Choose between processed markup and the exact post-unlock source. Default: false.
	ReturnPageSource *bool `json:"return_page_source" api:"return_page_source"`
	// Comma-separated list of fields to keep in the response, with dot notation for nested fields (e.g. name,location.city). Does not change the credit cost.
	Fields *string `json:"fields" api:"fields"`
	// Process the request asynchronously: the response is a job_id to poll on /v1/jobs/{id} (see waitForJob). Default: false.
	IsAsync *bool `json:"is_async" api:"is_async"`
	// Return HTTP 200 even when an error occurs; the error is then in the response body. Default: false.
	SoftFail *bool `json:"soft_fail" api:"soft_fail"`
}

func (p *WebsiteWebunlockerParams) params() []param {
	if p == nil {
		return nil
	}
	values := make([]param, 0, 6)
	if p.Query != "" {
		values = append(values, param{"query", p.Query})
	}
	if p.AllowRedirects != nil {
		values = append(values, param{"allow_redirects", *p.AllowRedirects})
	}
	if p.ReturnPageSource != nil {
		values = append(values, param{"return_page_source", *p.ReturnPageSource})
	}
	if p.Fields != nil {
		values = append(values, param{"fields", *p.Fields})
	}
	if p.IsAsync != nil {
		values = append(values, param{"is_async", *p.IsAsync})
	}
	if p.SoftFail != nil {
		values = append(values, param{"soft_fail", *p.SoftFail})
	}
	return values
}
