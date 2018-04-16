// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package domain

import (
	"testing"

	"github.com/seibert-media/k8s-ingress/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSyncer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

var _ = Describe("DomainFetcher", func() {
	var (
		domainFetcher *Fetcher
		httpClient    *mocks.DomainClient
	)

	BeforeEach(func() {
		httpClient = &mocks.DomainClient{}
		domainFetcher = &Fetcher{
			Client: httpClient,
		}
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
		It("does http call", func() {
			domainFetcher.Fetch()
			Expect(httpClient.DoCallCount()).To(Equal(1))
		})
		It("does not-nil request", func() {
			domainFetcher.Fetch()
			Expect(httpClient.DoArgsForCall(0)).NotTo(BeNil())
		})
	})
})
