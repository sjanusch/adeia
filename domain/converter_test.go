// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/k8s-ingress/model"
)

var _ = Describe("Fetcher", func() {
	var (
		domainConverter *Converter
		Domains         = []model.Domain{"http://server1.com/", "http://server2.com/"}
	)

	BeforeEach(func() {
		domainConverter = &Converter{
			Domains: Domains,
		}
	})

	Describe("Convert", func() {
		It("returns correct count of ingress objects", func() {
			ingress := domainConverter.Convert(Domains)
			Expect(ingress).ToNot(BeNil())
			Expect(ingress.Spec.Rules).To(HaveLen(2))
		})
		It("returns ingress objects with correct host", func() {
			ingress := domainConverter.Convert(Domains)
			Expect(ingress.Spec.Rules[0].Host).To(Equal("http://server1.com/"))
			Expect(ingress.Spec.Rules[1].Host).To(Equal("http://server2.com/"))
		})
	})
})
