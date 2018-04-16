package sync

import "github.com/seibert-media/k8s-ingress/model"

//go:generate counterfeiter -o ../mocks/ingress_fetcher.go --fake-name IngressFetcher . ingressFetcher
type ingressFetcher interface {
	Fetch() ([]model.Domain, error)
}

// IngressSyncer creates ingress for a list of domains
type IngressSyncer struct {
	Fetcher ingressFetcher
}

// Sync fetchs a list of domains an create ingresses
func (i *IngressSyncer) Sync() error {
	var _, err = i.Fetcher.Fetch()
	return err
}
