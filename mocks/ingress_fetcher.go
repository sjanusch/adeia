package mocks

import (
	"github.com/seibert-media/k8s-ingress/model"
)

type IngressFetcher struct {
	Counter int
	Error   error
}

func (i *IngressFetcher) Fetch() ([]model.Domain, error) {
	i.Counter++
	return nil, i.Error
}
