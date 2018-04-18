// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"io"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	"k8s.io/api/extensions/v1beta1"
)

var YamlMarshal = yaml.Marshal

// PrintApplier add ingress to k8sapplier/applier.go:18.
type PrintApplier struct {
	Out io.Writer
}

// Apply a list of domains
func (a *PrintApplier) Apply(ingress *v1beta1.Ingress) error {
	content, err := YamlMarshal(ingress)
	if err != nil {
		return errors.Wrap(err, "marshal yaml failed")
	}
	a.Out.Write(content)
	return nil
}
