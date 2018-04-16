package k8s_ingress

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/k8s-ingress/mocks"
)

func TestK8sIngress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8sIngress Suite")
}

var _ = Describe("K8sIngress", func() {
	var fetcher = &mocks.IngressFetcher{}
	var syncer = &ingressSyncer{
		fetcher: fetcher,
	}

	BeforeEach(func() {

	})

	Describe("Ingress Syncer", func() {
		It("calls ingress fetcher", func() {
			Expect(fetcher.Counter).To(Equal(0))
			Expect(syncer.Sync()).To(BeNil())
			Expect(fetcher.Counter).To(Equal(1))
		})
	})
})
