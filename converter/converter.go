package converter

import (
	"github.com/seibert-media/k8s-ingress/model"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type Converter struct {
	Domains []model.Domain
	Ingress v1beta1.Ingress

}

var (
	serverNamePtr = "test"
	serverPortPtr = "8080"
	namePtr = "GT"
	namespacePtr = "GT-NAMESPACE"
)

func (c *Converter) Convert() (v1beta1.Ingress, error) {


	var ingress = v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"kubernetes.io/ingress.class": "traefik",
			},
			Name:      namePtr,
			Namespace: namespacePtr,
		},
		Spec: v1beta1.IngressSpec{
			Rules: buildRuleSet(c.Domains),
		},
	}
	return ingress, nil
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
								ServiceName: serverNamePtr,
								ServicePort: intstr.Parse(serverPortPtr),
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