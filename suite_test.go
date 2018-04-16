// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestK8sIngress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8s Ingress Suite")
}
