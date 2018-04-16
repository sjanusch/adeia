// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import "github.com/seibert-media/k8s-ingress/model"

//go:generate counterfeiter -o ../mocks/ingress_fetcher.go --fake-name IngressFetcher . ingressFetcher
type ingressFetcher interface {
	Fetch() ([]model.Domain, error)
}

//go:generate counterfeiter -o ../mocks/ingress_applier.go --fake-name DomainApplier . domainApplier
type domainApplier interface {
	Apply([]model.Domain) error
}

// IngressSyncer creates ingress for a list of domains
type IngressSyncer struct {
	Fetcher ingressFetcher
	Applier domainApplier
}

// Sync fetchs a list of domains an create ingresses
func (i *IngressSyncer) Sync() error {
	list, err := i.Fetcher.Fetch()
	if err != nil {
		return err
	}
	return i.Applier.Apply(list)
}
