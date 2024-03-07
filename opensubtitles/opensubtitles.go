package opensubtitles

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
)

var (
	BaseURL   = "https://api.opensubtitles.com"
	UserAgent = "go-opensubtitles v1.0.0"
)

const (
	mediaTypeJSON = "application/json"

	headerContentType   = "Content-Type"
	headerAccept        = "Accept"
	headerUserAgent     = "User-Agent"
	headerApiKey        = "Api-Key"
	headerAuthorization = "Authorization"
)

// Credentials used to authenticate to make requests to the OpenSubtitles API.
type Credentials struct {
	Username string
	Password string
}

// A Client manages communication with the OpenSubtitles API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests.
	BaseURL *url.URL

	// Base URL for uploading files.
	UploadURL *url.URL

	// User agent used when communicating with the OpenSubtitles API.
	UserAgent string

	Credential Credentials

	rateMu sync.Mutex

	Token string

	ApiKey string
	//rateLimits [categories]Rate // Rate limits for the client as determined by the most recent API calls.

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the OpenSubtitles API.
	Authentication *AuthenticationService
	Discover       *DiscoverService
	Download       *DownloadService
	Find           *FindService
	Info           *InfoService

	User *UserData
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

// NewClient returns a new OpenSubtitles API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, token string, cred Credentials, apiKey string) (c *Client) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(BaseURL)
	//uploadURL, _ := url.Parse(uploadBaseURL)

	c = &Client{client: httpClient, BaseURL: baseURL, UserAgent: UserAgent, Token: token, ApiKey: apiKey}
	c.common.client = c
	c.Authentication = (*AuthenticationService)(&c.common)
	c.Discover = (*DiscoverService)(&c.common)
	c.Download = (*DownloadService)(&c.common)
	c.Find = (*FindService)(&c.common)
	c.Info = (*InfoService)(&c.common)

	//Check if struct Credential is not empty
	if (Credentials{}) != cred {
		c.Credential = cred
	}

	return
}

// Connect return a new Client with a working token by making the authentication
// with the Authentication Login function.
func (c *Client) Connect() (*Client, error) {
	if c == nil {
		return nil, errors.New("no client provided")
	}
	if len(c.Token) < 1 {
		login, resp, err := c.Authentication.Login(context.Background(), &c.Credential)
		if err == nil && login != nil {
			if len(login.Token) > 0 && resp.StatusCode == http.StatusOK {
				c.Token = login.Token
			} else {
				return nil, fmt.Errorf("Failed to get a token, got %q with status %q", login.Token, resp.Status)
			}
			if login.BaseURL != "" {
				newBaseURL, err := url.Parse(login.BaseURL)
				if err != nil {
					c.BaseURL = newBaseURL
				}
			}
			c.User = &login.User
		} else {
			return nil, errors.Join(err, errors.New("Failed to login"))
		}
	}
	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
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
	req.Header.Set(headerContentType, mediaTypeJSON)
	req.Header.Set(headerUserAgent, c.UserAgent)
	req.Header.Set(headerApiKey, c.ApiKey)
	if c.Token != "" {
		req.Header.Set(headerAuthorization, "Bearer "+c.Token)
	}
	//req.Header.Set("Accept", mediaTypeV3)
	return req, nil
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response, b []byte) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	if b != nil {
		_ = json.Unmarshal(b, errorResponse)
	}
	if errorResponse.Message != "" {
		return errorResponse
	}
	return nil
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
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (resp *http.Response, err error) {
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

	resp, err = c.client.Do(req)
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//log.Println("resp.Body:", string(b))
	//log.Println("resp.Status:", resp.Status)

	/*c.rateMu.Lock()
	c.rateLimits[rateLimitCategory] = response.Rate
	c.rateMu.Unlock()*/

	err = CheckResponse(resp, b)
	if err != nil {
		return nil, err
	}

	if v != nil { // This is true only of we received literally a nil, not some type with nil value.
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, bytes.NewReader(b))
		} else {
			decErr := json.NewDecoder(bytes.NewReader(b)).Decode(v)
			if decErr == io.EOF {
				log.Println("Ignoring decoding error:", decErr)
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}
