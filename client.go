package figi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const (
	ProdURL             = "https://api.openfigi.com"
	applicationTypeJson = "application/json"
)

type Client struct {
	logger  *slog.Logger
	c       *http.Client
	header  http.Header
	baseURL string
}

type ClientOpt func(*Client)

func WithHTTPClient(h *http.Client) ClientOpt      { return func(c *Client) { c.c = h } }
func WithLogHandler(h slog.Handler) ClientOpt      { return func(c *Client) { c.logger = slog.New(h) } }
func WithBaseURL(u string) ClientOpt               { return func(c *Client) { c.baseURL = u } }
func WithExtraHeader(k string, v string) ClientOpt { return func(c *Client) { c.header.Add(k, v) } }
func WithApiKey(key string) ClientOpt              { return WithExtraHeader("X-OPENFIGI-APIKEY", key) }

func New(opts ...ClientOpt) *Client {
	c := &Client{
		logger:  slog.New(slog.DiscardHandler),
		c:       http.DefaultClient,
		baseURL: ProdURL,
		header: http.Header{
			"Content-Type": []string{applicationTypeJson},
			"Accept":       []string{applicationTypeJson},
		},
	}

	for _, v := range opts {
		v(c)
	}

	return c
}

func (c *Client) req(ctx context.Context, meth, path string, body, target any) error {
	path = c.baseURL + "/" + path
	l := c.logger.With("path", path, "method", meth, "body", body)

	var bodyReader io.Reader
	if body != nil {
		buf, err := json.Marshal(body)
		if err != nil {
			l.ErrorContext(ctx, "failed marshaling request")
			return err
		}
		bodyReader = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, meth, path, bodyReader)
	if err != nil {
		l.ErrorContext(ctx, "failed creating request with context", "err", err)
		return err
	}
	req.Header = c.header.Clone()

	l = l.With("req", req)

	resp, err := c.c.Do(req)
	if err != nil {
		l.ErrorContext(ctx, "failed making request", "err", err)
		return err
	}
	l = l.With("statusCode", resp.StatusCode)

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		l.ErrorContext(ctx, "failed reading response body", "err", err)
		return err
	}
	l = l.With("resp", string(buf))

	if resp.StatusCode/100 != 2 {
		l.ErrorContext(ctx, "bad status code received")
		return fmt.Errorf("failed reading response %d:\n%s", resp.StatusCode, buf)
	}

	if err = json.Unmarshal(buf, target); err != nil {
		l.ErrorContext(ctx, "failed response unmarshal")
		return err
	}

	l.DebugContext(ctx, "request made")
	return nil
}
