package converter

import (
	"github.com/seibert-media/k8s-ingress/model"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Converter struct {
	Domains []model.Domain
}

var (
	serviceName = "test"
	serverPort  = "8080"
	name        = "GT"
	namespace   = "GT-NAMESPACE"
)

func (c *Converter) Convert() (*v1beta1.Ingress) {
	var ingress = v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "traefik",
			},
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1beta1.IngressSpec{
			Rules: buildRuleSet(c.Domains),
		},
	}
	return &ingress
}

func buildRuleSet(domains []model.Domain) ([]v1beta1.IngressRule) {
	var ingressRules []v1beta1.IngressRule
	for _, domain := range domains {
		ingressRule := v1beta1.IngressRule{
			Host: string(domain),
			IngressRuleValue: v1beta1.IngressRuleValue{
				HTTP: &v1beta1.HTTPIngressRuleValue{
					Paths: []v1beta1.HTTPIngressPath{
						{
							Path: "/",
							Backend: v1beta1.IngressBackend{
								ServiceName: serviceName,
								ServicePort: intstr.Parse(serverPort),
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
