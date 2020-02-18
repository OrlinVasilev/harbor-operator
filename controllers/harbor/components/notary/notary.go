package notary

import (
	"context"

	containerregistryv1alpha1 "github.com/ovh/harbor-operator/api/v1alpha1"
)

const (
	NotaryServerName = "notary-server"
	NotarySignerName = "notary-signer"
)

type Notary struct {
	harbor *containerregistryv1alpha1.Harbor
	Option Option
}

type Option struct {
	Priority *int32
}

func New(ctx context.Context, harbor *containerregistryv1alpha1.Harbor, opt Option) (*Notary, error) {
	return &Notary{
		harbor: harbor,
		Option: opt,
	}, nil
}