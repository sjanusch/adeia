// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import "github.com/seibert-media/k8s-ingress/model"

type fetcher struct{}

func (f *fetcher) Fetch() ([]model.Domain, error) {
	return []model.Domain{
		"www.example.com",
	}, nil
}
