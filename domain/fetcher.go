// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import "github.com/seibert-media/k8s-ingress/model"

// Fetcher get all domains.
type Fetcher struct{}

// Fetch domains from remote json endpoint.
func (f *Fetcher) Fetch() ([]model.Domain, error) {
	return []model.Domain{
		"www.example.com",
	}, nil
}
