package net

import (
	"github.com/Buccaneer69/wc-api-go/auth"
	"net/http"
)

// RequestEnricherMock ...
type RequestEnricherMock struct {
}

// EnrichRequest ...
func (a *RequestEnricherMock) EnrichRequest(r *http.Request, URL string) {
	auth := auth.Authenticator{}
	if auth.IsSsl(URL) {
		r.SetBasicAuth("key", "secret")
	}
}
