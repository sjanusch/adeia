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
	k8s_metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	k8s_discovery "k8s.io/client-go/discovery"
	k8s_dynamic "k8s.io/client-go/dynamic"
	restclient "k8s.io/client-go/rest"
)

var _ = Describe("K8sApplier", func() {
	var (
		k8sApplier    *ingress.K8sApplier
		testIngress   *k8s_v1beta1.Ingress
		ingressClient *mocks.IngressClient
	)

	BeforeEach(func() {
		k8sApplier = &ingress.K8sApplier{}
		testIngress = &k8s_v1beta1.Ingress{
			ObjectMeta: k8s_metav1.ObjectMeta{
				Annotations: map[string]string{
					"kubernetes.io/ingress.class": "traefik",
				},
				Name:      "test",
				Namespace: "testnamespace",
			},
			Spec: k8s_v1beta1.IngressSpec{
				Rules: []k8s_v1beta1.IngressRule{
					k8s_v1beta1.IngressRule{
						Host: string("example.com"),
						IngressRuleValue: k8s_v1beta1.IngressRuleValue{
							HTTP: &k8s_v1beta1.HTTPIngressRuleValue{
								Paths: []k8s_v1beta1.HTTPIngressPath{
									{
										Path: "/",
										Backend: k8s_v1beta1.IngressBackend{
											ServiceName: "test",
											ServicePort: intstr.Parse("80"),
										},
									},
								},
							},
						},
					},
				},
			},
		}
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
	})
})

func createK8sClients(cfg *restclient.Config) (*k8s_discovery.DiscoveryClient, k8s_dynamic.ClientPool, error) {
	discovery, err := k8s_discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return nil, nil, errors.Wrap(err, "creating k8s_discovery client failed")
	}
	dynamicPool := k8s_dynamic.NewDynamicClientPool(cfg)

	return discovery, dynamicPool, nil
}
