package questrade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/arianitu/go-questrade-oauth2"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
)

const (
	apiVersion = "v1"
)

type Client struct {
	ApiServerURL *url.URL
	Client       *http.Client
	IsPractice   bool
	UserAgent    string
}

type GeneralError struct {
	Response *http.Response

	// The Questrade API docs say that Code is a string, but they currently return a number
	Code    int
	Message string
}

func (r *GeneralError) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Code, r.Message)
}

func NewClient(refreshToken string, isPractice bool) (*Client, error) {
	conf := &questradeoauth2.Config{
		RefreshToken: refreshToken,
		IsPractice:   isPractice,
	}

	client, apiServer, err := conf.Client(oauth2.NoContext)
	if err != nil {
		return nil, err
	}

	apiServerURL, err := url.Parse(apiServer)
	if err != nil {
		return nil, err
	}

	return &Client{
		ApiServerURL: apiServerURL,
		Client:       client,
		IsPractice:   isPractice,
		UserAgent:    "",
	}, nil
}

// checks for errors in the HTTP response and parses them if there are any.
func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	generalError := &GeneralError{Response: r}
	err := json.NewDecoder(r.Body).Decode(generalError)
	if err != nil {
		return err
	}

	return generalError
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method string, urlStr string, body interface{}, v interface{}) error {
	var buf io.ReadWriter

	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}

	url, err := url.Parse(apiVersion + "/" + urlStr)
	if err != nil {
		return err
	}

	u := c.ApiServerURL.ResolveReference(url)
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return err
	}

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	err = checkResponse(resp)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}
