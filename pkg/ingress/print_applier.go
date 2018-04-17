// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"k8s.io/api/extensions/v1beta1"
	"fmt"
)

// PrintApplier add ingress to k8sapplier/applier.go:18.
type PrintApplier struct {
}

// Apply a list of domains
func (a *PrintApplier) Apply(ingress *v1beta1.Ingress) error {
	fmt.Printf("test")
	return nil
}
