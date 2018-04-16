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

	BeforeEach(func() {
		fetcher = &mocks.IngressFetcher{}
		syncer.Fetcher = fetcher
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
	})
})
