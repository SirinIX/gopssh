package direct

import "cmd-scaffold/pkg/http"

const (
	HeaderKeyAccept      = "Accept"
	HeaderKeyContentType = "Content-Type"

	HeaderValueAccept      = "application/json, text/plain, */*"
	HeaderValueContentType = "application/json;charset=UTF-8"

	CookieKeyDTStack   = "dtstack"
	CookieValueDTStack = "test"
)

type RequestBuilder struct {
	EmUrl      string
	HttpClient *http.Client
}

func NewRequestBuilder(emUrl string) (*RequestBuilder, error) {
	// Build http client
	httpClient, err := http.NewInsecureHTTPClient()
	if err != nil {
		return nil, err
	}

	return &RequestBuilder{
		EmUrl:      emUrl,
		HttpClient: httpClient,
	}, nil
}
