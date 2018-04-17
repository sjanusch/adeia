package ingress_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seibert-media/adeia/ingress"
	"k8s.io/api/extensions/v1beta1"
	"github.com/bborbe/io"
	"bytes"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("the PrintApplier", func() {
	var applier *ingress.PrintApplier
	var err error
	var out io.Writer
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
	It("returns no error", func() {
		applier.Apply(&v1beta1.Ingress{})
		Expect(out).To(gbytes.Say("test"))
	})
})
