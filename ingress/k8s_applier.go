// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"errors"

	k8s_v1beta1 "k8s.io/api/extensions/v1beta1"
)

//go:generate counterfeiter -o ../mocks/ingress_client.go --fake-name IngressClient . ingressClient
type ingressClient interface {
	Create(*k8s_v1beta1.Ingress) (*k8s_v1beta1.Ingress, error)
}

// K8sApplier add ingress to Client.
type K8sApplier struct {
	Client ingressClient
}

// Apply a list of domains.
func (a *K8sApplier) Apply(ingress *k8s_v1beta1.Ingress) error {
	if ingress == nil {
		return errors.New("ingress must not be nil")
	}
	if a.Client == nil {
		return errors.New("client must not be nil")
	}
	_, err := a.Client.Create(ingress)
	if err != nil {
		return err
	}
	return nil
}
