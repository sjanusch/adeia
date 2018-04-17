package converter

import (
	"github.com/seibert-media/k8s-ingress/model"
	"k8s.io/api/extensions/v1beta1"
)

type Converter struct {
	Domains []model.Domain
}

func (c *Converter) Convert() ([]v1beta1.Ingress, error) {
	return nil, nil
}