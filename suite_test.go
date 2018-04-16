// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/seibert-media/k8s-ingress/mocks"
	"github.com/seibert-media/k8s-ingress/sync"
)

func TestK8sIngress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8sIngress Suite")
}

var _ = Describe("K8s Ingress", func() {
	var fetcher *mocks.IngressFetcher
	var syncer = &sync.IngressSyncer{}
	var applier *mocks.DomainApplier

	BeforeEach(func() {
		fetcher = &mocks.IngressFetcher{}
		applier = &mocks.DomainApplier{}
		syncer.Fetcher = fetcher
		syncer.Applier = applier
	})

	Describe("Ingress Syncer", func() {
		It("calls ingress fetcher", func() {
			Expect(fetcher.FetchCallCount()).To(Equal(0))
			Expect(syncer.Sync()).To(BeNil())
			Expect(fetcher.FetchCallCount()).To(Equal(1))
		})
		It("return error when fetch fails", func() {
			fetcher.FetchReturns(nil, errors.New("Failed"))
			Expect(syncer.Sync()).NotTo(BeNil())
		})
		It("calls applier", func() {
			Expect(applier.ApplyCallCount()).To(Equal(0))
			syncer.Sync()
			Expect(applier.ApplyCallCount()).To(Equal(1))
		})
		It("does not apply if fetch fails", func() {
			fetcher.FetchReturns(nil, errors.New("Failed"))
			Expect(applier.ApplyCallCount()).To(Equal(0))
			syncer.Sync()
			Expect(applier.ApplyCallCount()).To(Equal(0))
		})
	})
})
