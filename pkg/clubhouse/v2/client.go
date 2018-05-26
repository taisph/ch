package v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	ClubHouseApiURL = "https://api.clubhouse.io/api/v2/"
)

var (
	UserAgent = "ch 0.0.0"
)

type Client struct {
	apiURL     *url.URL
	httpClient *http.Client
	token      string
	ua         string
}

func NewClient(token string) (*Client, error) {
	u, err := url.Parse(ClubHouseApiURL)
	if err != nil {
		return nil, err
	}
	return &Client{apiURL: u, httpClient: &http.Client{}, ua: UserAgent, token: token}, nil
}

func (c *Client) resolve(path string) (*url.URL, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	return c.apiURL.ResolveReference(u), nil
}

func (c *Client) get(path string, data interface{}, result interface{}) error {
	return c.dispatch(http.MethodGet, path, data, result)
}

func (c *Client) post(path string, data interface{}, result interface{}) error {
	return c.dispatch(http.MethodPost, path, data, result)
}

func (c *Client) dispatch(method string, path string, data interface{}, result interface{}) error {
	u, err := c.resolve(path)
	if err != nil {
		return err
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	v := req.URL.Query()
	v.Set("token", c.token)
	req.URL.RawQuery = v.Encode()

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.ua)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode >= 400 {
		type apiError struct {
			Msg string `json:"message"`
		}
		var e apiError
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&e); err != nil {
			return fmt.Errorf("error sending request and decoding response: %s %s: %d: %s", method, u.String(), res.StatusCode, err.Error())
		}
		return fmt.Errorf("error sending request: %s %s: %d: %s", method, u.String(), res.StatusCode, e.Msg)
	}
	if err := decoder.Decode(result); err != nil {
		return fmt.Errorf("error decoding response while sending request: %s %s: %s", method, u.String(), err.Error())
	}

	return nil
}
