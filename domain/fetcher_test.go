// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSyncer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

var _ = Describe("DomainFetcher", func() {
	var (
		domainFetcher *fetcher
	)

	BeforeEach(func() {
		domainFetcher = &fetcher{}
	})

	Describe("Fetcher", func() {
		It("returns no error", func() {
			_, err := domainFetcher.Fetch()
			Expect(err).To(BeNil())
		})
		It("returns one domain", func() {
			list, _ := domainFetcher.Fetch()
			Expect(list).To(HaveLen(1))
		})
	})
})
