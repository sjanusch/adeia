// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"github.com/pkg/errors"
	"github.com/seibert-media/k8s-ingress/pkg/domain"
	"k8s.io/api/extensions/v1beta1"
)

//go:generate counterfeiter -o mocks/ingress_fetcher.go --fake-name IngressFetcher . fetcher
type fetcher interface {
	Fetch() ([]domain.Domain, error)
}

//go:generate counterfeiter -o mocks/ingress_applier.go --fake-name IngressApplier . applier
type applier interface {
	Apply(ingress *v1beta1.Ingress) error
}

//go:generate counterfeiter -o 	mocks/ingress_creator.go --fake-name IngressCreator . creator
type creator interface {
	Create(domains []domain.Domain) *v1beta1.Ingress
}

// Syncer creates ingress for a list of domains
type Syncer struct {
	Fetcher fetcher
	Creator creator
	Applier applier
}

// Sync fetches a list of domains an create ingresses
func (i *Syncer) Sync() error {
	domains, err := i.Fetcher.Fetch()
	if err != nil {
		return errors.Wrap(err, "fetch domain failed")
	}
	return i.Applier.Apply(i.Creator.Create(domains))
}
