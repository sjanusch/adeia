// Copyright 2018 The adeia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/seibert-media/adeia/ingress"
	"github.com/seibert-media/adeia/mocks"
	k8s_v1beta1 "k8s.io/api/extensions/v1beta1"
)

var _ = Describe("K8sApplier", func() {
	var (
		k8sApplier    *ingress.K8sApplier
		testIngress   *k8s_v1beta1.Ingress
		ingressClient *mocks.IngressClient
	)

	BeforeEach(func() {
		k8sApplier = &ingress.K8sApplier{}
		testIngress = &k8s_v1beta1.Ingress{}
		ingressClient = &mocks.IngressClient{}
	})

	Describe("Apply", func() {
		It("returns error on nil ingress", func() {
			Expect(k8sApplier.Apply(nil)).NotTo(BeNil())
		})
		It("returns error on nil client", func() {
			Expect(k8sApplier.Apply(testIngress)).NotTo(BeNil())
		})
		It("returns error when creating ingress fails", func() {
			k8sApplier.Client = ingressClient
			ingressClient.CreateReturns(nil, errors.New("test"))
			Expect(k8sApplier.Apply(testIngress)).NotTo(BeNil())
		})
		It("does not return error when creating ingress is successful", func() {
			k8sApplier.Client = ingressClient
			ingressClient.CreateReturns(nil, nil)
			Expect(k8sApplier.Apply(testIngress)).To(BeNil())
		})
	})
})
