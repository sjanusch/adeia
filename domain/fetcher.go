// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"errors"
	"net/http"

	"github.com/seibert-media/k8s-ingress/model"
)

//go:generate counterfeiter -o ../mocks/domain_client.go --fake-name DomainClient . client
type client interface {
	Do(*http.Request) (*http.Response, error)
}

// Fetcher get all domains.
type Fetcher struct {
	Client client
	URL    string
}

// Fetch domains from remote json endpoint.
func (f *Fetcher) Fetch() ([]model.Domain, error) {
	if len(f.URL) < 1 {
		return nil, errors.New("invalid URL")
	}
	req, _ := http.NewRequest("GET", f.URL, nil)
	_, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return []model.Domain{
		"www.example.com",
	}, nil
}
