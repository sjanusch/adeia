package ingress

import (
	"testing"

	"github.com/seibert-media/k8s-ingress/mocks"
	"github.com/seibert-media/k8s-ingress/model"
)

func TestSyncerSync(t *testing.T) {
	f := &mocks.IngressFetcher{
		FetchStub: func() ([]model.Domain, error) {
			return nil, nil
		},
	}
	s := &Syncer{
		Fetcher: f,
	}
	if err := s.Sync(); err != nil {
		t.Error("error unexpected", err)
	}
}
