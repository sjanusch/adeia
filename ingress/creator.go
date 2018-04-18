// Copyright 2018 The K8s-Ingress Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ingress

import (
	"github.com/seibert-media/adeia/domain"
	k8s_v1beta1 "k8s.io/api/extensions/v1beta1"
	k8s_metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s_intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// Creator for transform domain to ingress
type Creator struct {
	Ingressname string
	Servicename string
	Serviceport string
	Namespace   string
}

// Create Ingress for the given domains.
func (c *Creator) Create(domains []domain.Domain) *k8s_v1beta1.Ingress {
	return &k8s_v1beta1.Ingress{
		TypeMeta: k8s_metav1.TypeMeta{
			APIVersion: "extensions/v1beta1",
			Kind:       "Ingress",
		},
		ObjectMeta: k8s_metav1.ObjectMeta{
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "traefik",
			},
			Name:      c.Ingressname,
			Namespace: c.Namespace,
		},
		Spec: k8s_v1beta1.IngressSpec{
			Rules: c.buildRuleSet(domains),
		},
	}
}

func (c *Creator) buildRuleSet(domains []domain.Domain) []k8s_v1beta1.IngressRule {
	var ingressRules []k8s_v1beta1.IngressRule
	for _, domain := range domains {
		ingressRule := k8s_v1beta1.IngressRule{
			Host: string(domain),
			IngressRuleValue: k8s_v1beta1.IngressRuleValue{
				HTTP: &k8s_v1beta1.HTTPIngressRuleValue{
					Paths: []k8s_v1beta1.HTTPIngressPath{
						{
							Path: "/",
							Backend: k8s_v1beta1.IngressBackend{
								ServiceName: c.Servicename,
								ServicePort: k8s_intstr.Parse(c.Serviceport),
							},
						},
					},
				},
			},
		}
		ingressRules = append(ingressRules, ingressRule)
	}
	return ingressRules
}
