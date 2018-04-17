package applier

import (
	"github.com/seibert-media/k8s-ingress/model"
	"k8s.io/api/extensions/v1beta1"
)

//go:generate counterfeiter -o ../mocks/ingress_converter.go --fake-name DomainConverter . converter
type converter interface {
	Convert([]model.Domain) *v1beta1.Ingress
}

type Applier struct {
	Converter converter
}

func (a *Applier) Apply(domains []model.Domain) error {
	return nil
}
