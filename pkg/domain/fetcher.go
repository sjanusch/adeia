// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

//go:generate counterfeiter -o ../mocks/domain_client.go --fake-name DomainClient . client
type client interface {
	Get(string) (*http.Response, error)
}

// Fetcher get all domains.
type Fetcher struct {
	Client client
	URL    string
}

// Fetch domains from remote json endpoint.
func (f *Fetcher) Fetch() ([]Domain, error) {
	if len(f.URL) < 1 {
		return nil, errors.New("invalid URL")
	}
	resp, err := f.Client.Get(f.URL)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("received empty response")
	}
	result := []Domain{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "parse json failed")
	}
	return result, nil
}
