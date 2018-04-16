// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import "github.com/seibert-media/k8s-ingress/model"

//go:generate counterfeiter -o ../mocks/ingress_fetcher.go --fake-name IngressFetcher . fetcher
type fetcher interface {
	Fetch() ([]model.Domain, error)
}

//go:generate counterfeiter -o ../mocks/ingress_applier.go --fake-name DomainApplier . applier
type applier interface {
	Apply([]model.Domain) error
}

// Syncer creates ingress for a list of domains
type Syncer struct {
	Fetcher fetcher
	Applier applier
}

// Sync fetches a list of domains an create ingresses
func (i *Syncer) Sync() error {
	list, err := i.Fetcher.Fetch()
	if err != nil {
		return err
	}
	return i.Applier.Apply(list)
}
