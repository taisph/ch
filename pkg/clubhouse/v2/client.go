/*
 * Copyright 2018 Tais P. Hansen
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

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
	UserAgent = "ch 0.0.1"
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
			Msg    string      `json:"message"`
			Errors interface{} `json:"errors"`
			Tag    string      `json:"tag"`
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
