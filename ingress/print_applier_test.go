// Copyright 2018 The adeia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"bytes"
	"errors"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	k8s_v1beta1 "k8s.io/api/extensions/v1beta1"
)

var _ = Describe("the PrintApplier", func() {
	var applier *PrintApplier
	var err error
	var out *bytes.Buffer
	BeforeEach(func() {
		yamlMarshal = yaml.Marshal
		out = &bytes.Buffer{}
		applier = &PrintApplier{
			Out: out,
		}
	})
	It("returns no error", func() {
		err = applier.Apply(&k8s_v1beta1.Ingress{})
		Expect(err).To(BeNil())
	})
	It("returns error if marshal fails", func() {
		yamlMarshal = func(interface{}) ([]byte, error) { return nil, errors.New("banana") }
		err = applier.Apply(&k8s_v1beta1.Ingress{})
		Expect(err).NotTo(BeNil())
		Expect(err.Error()).To(Equal("marshal yaml failed: banana"))
	})
	It("returns content length greater zero", func() {
		applier.Apply(&k8s_v1beta1.Ingress{})
		Expect(out.Len()).Should(BeNumerically(">", 0))
	})
	It("returns serialized ingress", func() {
		applier.Apply(&k8s_v1beta1.Ingress{})
		Expect(out.String()).To(Equal(`metadata:
  creationTimestamp: null
spec: {}
status:
  loadBalancer: {}
`))
	})
})
