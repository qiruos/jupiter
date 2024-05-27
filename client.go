package jupiter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/qiruos/jupiter/utils"
)

const (
	// ContentTypeJSON is the content type for JSON.
	ContentTypeJSON = "application/json"
)

type (
	// Client is a Jupiter client that can be used to make requests to the Jupiter API.
	Client struct {
		client *http.Client

		apiURL                   string
		endpointQuote            string
		endpointSwap             string
		endpointSwapInstructions string
	}

	// ClientOption is a function that can be used to configure a Jupiter client.
	ClientOption func(*Client)

	// Response is a generic response structure.
	Response struct {
		Data        json.RawMessage `json:"data"`
		TimeTaken   float64         `json:"timeTaken"`
		ContextSlot int64           `json:"contextSlot"`
	}
)

// NewClient returns a new Jupiter client.
func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},

		apiURL:                   "https://quote-api.jup.ag/v6",
		endpointQuote:            "/quote",
		endpointSwap:             "/swap",
		endpointSwapInstructions: "/swap-instructions",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// get makes a GET request to the specified endpoint with the given parameters.
// It returns the response as is without parsing or any error encountered.
// The caller is responsible for closing the response body.
func (c *Client) get(endpoint string, params interface{}) (*http.Response, error) {
	uv, err := utils.StructToUrlValues(params)
	if err != nil {
		return nil, fmt.Errorf("failed to convert params to url values: %w", err)
	}

	parsedURL, err := url.Parse(c.apiURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if len(uv) > 0 {
		parsedURL.RawQuery = uv.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}
	req.Header.Set("Accept", ContentTypeJSON)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}

	return resp, nil
}

// postRaw makes a POST request to the specified URL with the given parameters.
// It returns the response as is without parsing or any error encountered.
// The caller is responsible for closing the response body.
func (c *Client) post(endpoint string, params interface{}) (*http.Response, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal POST params: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.apiURL+endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %w", err)
	}
	req.Header.Set("Content-Type", ContentTypeJSON)
	req.Header.Set("Accept", ContentTypeJSON)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %w", err)
	}

	return resp, nil
}

// Quote returns a quote for a given input mint, output mint and amount
func (c *Client) Quote(params QuoteParams) (*QuoteResponse, error) {
	resp, err := c.get(c.endpointQuote, params)
	if err != nil {
		return nil, fmt.Errorf("failed to make quote request: %w", err)
	}

	buf, err := io.ReadAll(resp.Body)

	var quotes QuoteResponse
	if err := json.Unmarshal(buf, &quotes); err != nil {
		return nil, fmt.Errorf("failed to parse quote response: %w", err)
	}

	return &quotes, nil
}

// Swap returns swap base64 serialized transaction for a route.
// The caller is responsible for signing the transactions.
func (c *Client) Swap(params SwapParams) (string, error) {
	resp, err := c.post(c.endpointSwap, params)
	if err != nil {
		return "", fmt.Errorf("failed to make swap request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response SwapResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return response.SwapTransaction, nil
}

// SwapInstructions Returns instructions that you can use from the quote you get from /quote.
// The caller is responsible for signing the transactions.
func (c *Client) SwapInstructions(params SwapParams) (*SwapInstructionsResp, error) {
	resp, err := c.post(c.endpointSwapInstructions, params)
	if err != nil {
		return nil, fmt.Errorf("failed to make swap request: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	var response SwapInstructionsResp
	if err := json.Unmarshal(buf, &response); err != nil {
		return nil, fmt.Errorf("failed to parse quote response: %w", err)
	}

	return &response, nil
}
