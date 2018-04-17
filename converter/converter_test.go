// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package converter

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/k8s-ingress/model"
)

func TestSyncer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

var _ = Describe("Fetcher", func() {
	var (
		domainConverter *Converter
		Domains =  []model.Domain{"http://server.com/domains", "http://server.com/domains"}
	)

	BeforeEach(func() {
		domainConverter = &Converter{
			Domains : Domains,
			}
	})

	Describe("Convert", func() {
		It("returns no error", func() {
			_, err := domainConverter.Convert()
			Expect(err).To(BeNil())
		})
		It("returns correct count of ingress objects", func() {
			ingresses, _ := domainConverter.Convert()
			Expect(ingresses).To(HaveLen(1))
		})
	})
})
