// Copyright 2018 The adeia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress_test

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/adeia/ingress"
	"k8s.io/api/extensions/v1beta1"
)

var _ = Describe("the PrintApplier", func() {
	var applier *ingress.PrintApplier
	var err error
	var out *bytes.Buffer
	BeforeEach(func() {
		out = &bytes.Buffer{}
		applier = &ingress.PrintApplier{
			Out: out,
		}

	})

	It("returns no error", func() {
		err = applier.Apply(&v1beta1.Ingress{})
		Expect(err).To(BeNil())
	})
	It("returns content length greater zero", func() {
		applier.Apply(&v1beta1.Ingress{})
		Expect(out.Len()).Should(BeNumerically(">", 0))
	})
	It("returns serialized ingress", func() {
		applier.Apply(&v1beta1.Ingress{})
		Expect(out.String()).To(Equal(`metadata:
  creationTimestamp: null
spec: {}
status:
  loadBalancer: {}
`))
	})
})
