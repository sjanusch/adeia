package ingress

import (
	"k8s.io/api/extensions/v1beta1"
)

// Applier add ingress to k8sapplier/applier.go:18.
type Applier struct {
}

// Apply a list of domains
func (a *Applier) Apply(ingress *v1beta1.Ingress) error {
	return nil
}
