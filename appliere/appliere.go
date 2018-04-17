package appliere

import (
	"fmt"
	"github.com/seibert-media/k8s-ingress/model"
)

type Appliere struct{}

func (a *Appliere) Apply([]model.Domain) error {
	fmt.Println("Ingressobject")
	return nil
}
