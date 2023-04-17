// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adeia

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"github.com/seibert-media/adeia/domain"
	k8s_networkingv1 "k8s.io/api/networking/v1"
)

//go:generate counterfeiter -o mocks/ingress_fetcher.go --fake-name IngressFetcher . fetcher
type fetcher interface {
	Fetch() ([]domain.Domain, error)
}

//go:generate counterfeiter -o mocks/ingress_applier.go --fake-name IngressApplier . applier
type applier interface {
	Apply(ingress *k8s_networkingv1.Ingress) error
}

//go:generate counterfeiter -o 	mocks/ingress_creator.go --fake-name IngressCreator . creator
type creator interface {
	Create(domains []domain.Domain) *k8s_networkingv1.Ingress
}

// Syncer creates Ingress for a list of domains
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
	glog.V(2).Infof("fetched %d domains", len(domains))
	glog.V(4).Infof("domains = %v", domains)
	return i.Applier.Apply(i.Creator.Create(domains))
}
