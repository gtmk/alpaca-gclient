package alpacaio

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	ej "github.com/mailru/easyjson"
)

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Client Configuration                     ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

const (
	DateLayoutISO = "2006-01-02"
	apiURL        = "https://data.alpaca.markets"
)

type Client struct {
	baseURL    string
	apiKey     string
	secretKey  string
	httpClient *http.Client
}

type Error struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), e.Message)
}

func NewClient(apiKey, secretKey string, options ...func(*Client)) *Client {
	client := &Client{
		apiKey:     apiKey,
		secretKey:  secretKey,
		httpClient: &http.Client{},
	}

	// apply options
	for _, option := range options {
		option(client)
	}

	// set default values
	if client.baseURL == "" {
		client.baseURL = apiURL
	}
	return client
}

func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) func(*Client) {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

//func (c *Client) GetJSON(ctx context.Context, endpoint string, v interface{}) error {
//	//address, err := c.addToken(endpoint)
//	//if err != nil {
//	//	return err
//	//}
//	data, err := c.getBytes(ctx, address)
//	if err != nil {
//		return err
//	}
//	return json.Unmarshal(data, v)
//}

func (c *Client) GetBytes(ctx context.Context, endpoint string) ([]byte, error) {
	//address, err := c.addToken(endpoint)
	//if err != nil {
	//	return []byte{}, err
	//}
	return c.getBytes(ctx, endpoint)
}

func (c *Client) getBytes(ctx context.Context, address string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.baseURL+address, nil)
	if err != nil {
		return []byte{}, err
	}
	c.addHeaders(req)
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		msg := ""

		if err == nil {
			msg = string(b)
		}

		return []byte{}, Error{StatusCode: resp.StatusCode, Message: msg}
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) addHeaders(req *http.Request) {
	req.Header.Add("APCA-API-KEY-ID", c.apiKey)
	req.Header.Add("APCA-API-SECRET-KEY", c.secretKey)
}

//func (c *Client) addToken(endpoint string) (string, error) {
//	u, err := url.Parse(c.baseURL + endpoint)
//	if err != nil {
//		return "", err
//	}
//	v := u.Query()
//	v.Add("apiKey", c.token)
//	u.RawQuery = v.Encode()
//	return u.String(), nil
//}

func (c *Client) endpointWithOpts(endpoint string, opts *RequestOptions) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}
	return endpoint, nil
}

func (c *Client) StockAggregates(ticker string, timespan Timespan, from, to time.Time, opts *RequestOptions) (*StockBarsResponse, error) {
	if opts == nil {
		opts = &RequestOptions{Limit: 10000}
	}
	opts.Start = from.Format(DateLayoutISO)
	opts.End = to.Format(DateLayoutISO)
	opts.Timeframe = string(timespan)
	var out StockBarsResponse
	endpoint := fmt.Sprintf("/v2/stocks/%s/bars", url.PathEscape(ticker))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	bts, err := c.getBytes(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}
	err = ej.Unmarshal(bts, &out)
	return &out, err
}

func (c *Client) StockTrades(ticker string, from, to time.Time, opts *RequestOptions) (*StockTradesResponse, error) {
	if opts == nil {
		opts = &RequestOptions{Limit: 10000}
	}
	opts.Start = from.Format(DateLayoutISO)
	opts.End = to.Format(DateLayoutISO)
	var out StockTradesResponse
	endpoint := fmt.Sprintf(`/v2/stocks/%s/trades`, url.PathEscape(ticker))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	bts, err := c.getBytes(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}
	err = ej.Unmarshal(bts, &out)
	return &out, err
}

func (c *Client) StockQuotes(ticker string, from, to time.Time, opts *RequestOptions) (*StockQuotesResponse, error) {
	if opts == nil {
		opts = &RequestOptions{Limit: 10000}
	}
	opts.Start = from.Format(DateLayoutISO)
	opts.End = to.Format(DateLayoutISO)
	var out StockQuotesResponse
	endpoint := fmt.Sprintf("/v2/stocks/%s/quotes", url.PathEscape(ticker))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	bts, err := c.getBytes(context.Background(), endpoint)
	if err != nil {
		return nil, err
	}
	err = ej.Unmarshal(bts, &out)
	return &out, err
}