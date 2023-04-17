// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"io"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	k8s_networkingv1 "k8s.io/api/networking/v1"
)

var yamlMarshal = yaml.Marshal

// PrintApplier add Ingress to k8sapplier/applier.go:18.
type PrintApplier struct {
	Out io.Writer
}

// Apply a list of domains
func (a *PrintApplier) Apply(ingress *k8s_networkingv1.Ingress) error {
	content, err := yamlMarshal(ingress)
	if err != nil {
		return errors.Wrap(err, "marshal yaml failed")
	}
	a.Out.Write(content)
	return nil
}
