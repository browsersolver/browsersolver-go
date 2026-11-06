// Package browsersolver is the official Go SDK for the BrowserSolver API (https://www.browsersolver.com).
//
// Unlock, render, crawl and screenshot any website — anti-bot bypass and technology detection through one API.
//
// Create one client and reuse it:
//
//	client, err := browsersolver.NewClient("", nil) // reads BROWSERSOLVER_API_KEY
//	response, err := client.WebsiteWebunlocker(ctx, ...)
//
// Every documented endpoint has a typed method; Client.Request reaches any other path.
package browsersolver
