// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "rankdb": elements Resource Client
//
// Command:
// $ goagen
// --design=github.com/Vivino/rankdb/api/design
// --out=$(GOPATH)/src/github.com/Vivino/rankdb/api
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CreateElementsPath computes a request path to the create action of elements.
func CreateElementsPath(listID string) string {
	param0 := listID

	return fmt.Sprintf("/lists/%s/elements", param0)
}

// Create Element in list
func (c *Client) CreateElements(ctx context.Context, path string, payload *Element, range_ *int, contentType string) (*http.Response, error) {
	req, err := c.NewCreateElementsRequest(ctx, path, payload, range_, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateElementsRequest create the request corresponding to the create action endpoint of the elements resource.
func (c *Client) NewCreateElementsRequest(ctx context.Context, path string, payload *Element, range_ *int, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if range_ != nil {
		tmp49 := strconv.Itoa(*range_)
		values.Set("range", tmp49)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteElementsPath computes a request path to the delete action of elements.
func DeleteElementsPath(listID string, elementID string) string {
	param0 := listID
	param1 := elementID

	return fmt.Sprintf("/lists/%s/elements/%s", param0, param1)
}

// Delete element in list. If element is not found the operation is considered a success
func (c *Client) DeleteElements(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteElementsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteElementsRequest create the request corresponding to the delete action endpoint of the elements resource.
func (c *Client) NewDeleteElementsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeleteMultiElementsPath computes a request path to the delete-multi action of elements.
func DeleteMultiElementsPath(listID string) string {
	param0 := listID

	return fmt.Sprintf("/lists/%s/elements", param0)
}

// Delete Multiple Elements in list.If an element does not exist, success is returned.
//
func (c *Client) DeleteMultiElements(ctx context.Context, path string, elementIds []string) (*http.Response, error) {
	req, err := c.NewDeleteMultiElementsRequest(ctx, path, elementIds)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteMultiElementsRequest create the request corresponding to the delete-multi action endpoint of the elements resource.
func (c *Client) NewDeleteMultiElementsRequest(ctx context.Context, path string, elementIds []string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	for _, p := range elementIds {
		tmp50 := p
		values.Add("element_ids", tmp50)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetElementsPath computes a request path to the get action of elements.
func GetElementsPath(listID string, elementID string) string {
	param0 := listID
	param1 := elementID

	return fmt.Sprintf("/lists/%s/elements/%s", param0, param1)
}

// Get Element in list
func (c *Client) GetElements(ctx context.Context, path string, range_ *int) (*http.Response, error) {
	req, err := c.NewGetElementsRequest(ctx, path, range_)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetElementsRequest create the request corresponding to the get action endpoint of the elements resource.
func (c *Client) NewGetElementsRequest(ctx context.Context, path string, range_ *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if range_ != nil {
		tmp51 := strconv.Itoa(*range_)
		values.Set("range", tmp51)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetAroundElementsPath computes a request path to the get-around action of elements.
func GetAroundElementsPath(listID string, elementID string) string {
	param0 := listID
	param1 := elementID

	return fmt.Sprintf("/lists/%s/elements/%s/around", param0, param1)
}

// Get relation of one element to multiple specific elements.
// The element will have local_from_top and local_from_bottom populated.Elements that are not found are ignored.
//
func (c *Client) GetAroundElements(ctx context.Context, path string, payload *MultiElement, range_ *int, contentType string) (*http.Response, error) {
	req, err := c.NewGetAroundElementsRequest(ctx, path, payload, range_, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAroundElementsRequest create the request corresponding to the get-around action endpoint of the elements resource.
func (c *Client) NewGetAroundElementsRequest(ctx context.Context, path string, payload *MultiElement, range_ *int, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if range_ != nil {
		tmp52 := strconv.Itoa(*range_)
		values.Set("range", tmp52)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetMultiElementsPath computes a request path to the get-multi action of elements.
func GetMultiElementsPath(listID string) string {
	param0 := listID

	return fmt.Sprintf("/lists/%s/elements/find", param0)
}

// Get Multiple Elements in list.
// Will return 404 if list cannot be found, OK even if no elements are found.
func (c *Client) GetMultiElements(ctx context.Context, path string, payload *MultiElement, contentType string) (*http.Response, error) {
	req, err := c.NewGetMultiElementsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetMultiElementsRequest create the request corresponding to the get-multi action endpoint of the elements resource.
func (c *Client) NewGetMultiElementsRequest(ctx context.Context, path string, payload *MultiElement, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// PutElementsPath computes a request path to the put action of elements.
func PutElementsPath(listID string, elementID string) string {
	param0 := listID
	param1 := elementID

	return fmt.Sprintf("/lists/%s/elements/%s", param0, param1)
}

// Update element in list
// If element does not exist, it is created in list.
// Element ID in payload an url must match.
func (c *Client) PutElements(ctx context.Context, path string, payload *Element, range_ *int, contentType string) (*http.Response, error) {
	req, err := c.NewPutElementsRequest(ctx, path, payload, range_, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPutElementsRequest create the request corresponding to the put action endpoint of the elements resource.
func (c *Client) NewPutElementsRequest(ctx context.Context, path string, payload *Element, range_ *int, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if range_ != nil {
		tmp53 := strconv.Itoa(*range_)
		values.Set("range", tmp53)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// PutMultiElementsPayload is the elements put-multi action payload.
type PutMultiElementsPayload []*Element

// PutMultiElementsPath computes a request path to the put-multi action of elements.
func PutMultiElementsPath(listID string) string {
	param0 := listID

	return fmt.Sprintf("/lists/%s/elements", param0)
}

// Update Multiple Elements in list.If element does not exist, it is created in list.
// The returned "not_found" field will never be preset.
func (c *Client) PutMultiElements(ctx context.Context, path string, payload PutMultiElementsPayload, results *bool, contentType string) (*http.Response, error) {
	req, err := c.NewPutMultiElementsRequest(ctx, path, payload, results, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPutMultiElementsRequest create the request corresponding to the put-multi action endpoint of the elements resource.
func (c *Client) NewPutMultiElementsRequest(ctx context.Context, path string, payload PutMultiElementsPayload, results *bool, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if results != nil {
		tmp54 := strconv.FormatBool(*results)
		values.Set("results", tmp54)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
