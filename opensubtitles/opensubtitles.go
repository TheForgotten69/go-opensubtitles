package opensubtitles

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/go-querystring/query"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
)

const (
	defaultBaseURL = "https://www.opensubtitles.com"
	userAgent      = "go-opensubtitles"

	mediaTypeJSON = "application/json"
	mediaTypeForm = "application/x-www-form-urlencoded"

	headerContentType = "Content-Type"
	headerAccept      = "Accept"
	headerUserAgent   = "User-Agent"
)

// Credentials used to authenticate to make requests to the OpenSubtitles API.
type Credentials struct {
	Username string
	Password string
}

// A Client manages communication with the GitHub API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	// Base URL for uploading files.
	UploadURL *url.URL

	// User agent used when communicating with the GitHub API.
	UserAgent string

	Credential Credentials

	rateMu sync.Mutex

	Token string
	//rateLimits [categories]Rate // Rate limits for the client as determined by the most recent API calls.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the GitHub API.
	Authentication *AuthenticationService
	Discover       *DiscoverService
	Download       *DownloadService
	Find           *FindService
	Info           *InfoService
	Search         *SearchService
}

type service struct {
	client *Client
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns a new GitHub API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, token string, cred Credentials) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	//uploadURL, _ := url.Parse(uploadBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent, Token: token}
	c.common.client = c
	c.Authentication = (*AuthenticationService)(&c.common)
	c.Discover = (*DiscoverService)(&c.common)
	c.Download = (*DownloadService)(&c.common)
	c.Find = (*FindService)(&c.common)
	c.Info = (*InfoService)(&c.common)
	c.Search = (*SearchService)(&c.common)

	//Check if struct Credential is not empty
	if (Credentials{}) != cred {
		c.Credential = cred
	}
	if len(c.Token) < 1 {
		log, resp, _ := c.Authentication.Login(context.Background(), &c.Credential)
		if (&LoggedIn{}) != log {
			if len(log.Token) > 0 && resp.Status == "200" {
				c.Token = log.Token
			} else {
				return nil, errors.New("Wrong Username/Password")
			}

		}
	}

	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body string) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	/*var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		/*enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}*/
	var b io.Reader
	if body != "" {
		b = strings.NewReader(body)
	}
	req, err := http.NewRequest(method, u.String(), b)
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerAccept, mediaTypeJSON)
	req.Header.Set("Content-Type", mediaTypeJSON)
	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}
	//req.Header.Set("Accept", mediaTypeV3)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)

	//rateLimitCategory := category(req.URL.Path)

	// If we've hit rate limit, don't make further requests before Reset time.
	/*if err := c.checkRateLimitBeforeDo(req, rateLimitCategory); err != nil {
		return &Response{
			Response: err.Response,
			Rate:     err.Rate,
		}, err
	}*/

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}

	defer resp.Body.Close()

	/*b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}*/

	c.rateMu.Lock()
	//c.rateLimits[rateLimitCategory] = response.Rate
	c.rateMu.Unlock()

	//err = CheckResponse(resp, b)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	// Block is only populated on certain types of errors such as code 451.
	// See https://developer.github.com/changes/2016-03-17-the-451-status-code-is-now-supported/
	// for more information.
	Block *struct {
		Reason string `json:"reason,omitempty"`
	} `json:"block,omitempty"`
	// Most errors will also include a documentation_url field pointing
	// to some content that might help you resolve the error, see
	// https://developer.github.com/v3/#client-errors
	DocumentationURL string `json:"documentation_url,omitempty"`
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have response
// body, and a JSON response body that maps to ErrorResponse.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
// *AcceptedError for 202 Accepted status codes,
// and *TwoFactorAuthError for two-factor authentication errors.
func CheckResponse(r *http.Response, b []byte) error {
	errorResponse := &ErrorResponse{Response: r}
	if b != nil {
		json.Unmarshal(b, errorResponse)
	}
	return nil

}
