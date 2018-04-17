// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"k8s.io/api/extensions/v1beta1"
)

// K8sApplier add ingress to k8sapplier/applier.go:18.
type K8sApplier struct {
}

// Apply a list of domains
func (a *K8sApplier) Apply(ingress *v1beta1.Ingress) error {
	panic("implement me")
}
