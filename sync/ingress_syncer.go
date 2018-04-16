package sync

import "github.com/seibert-media/k8s-ingress/model"

//go:generate counterfeiter -o mocks/ingress_fetcher.go --fake-name IngressFetcher . ingressFetcher
type ingressFetcher interface {
	Fetch() ([]model.Domain, error)
}

type IngressSyncer struct {
	Fetcher ingressFetcher
}

func (i *IngressSyncer) Sync() error {
	var _, err = i.Fetcher.Fetch()
	return err
}
