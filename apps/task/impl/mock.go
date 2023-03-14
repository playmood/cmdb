package impl

import (
	"context"
	"github.com/playmood/cmdb/apps/secret"
)

type secretMock struct {
	secret.UnimplementedServiceServer
}

func (m *secretMock) CreateSecret(context.Context, *secret.CreateSecretRequest) (*secret.Secret, error) {
	return nil, nil
}

func (m *secretMock) QuerySecret(context.Context, *secret.QuerySecretRequest) (*secret.SecretSet, error) {
	return nil, nil
}

func (m *secretMock) DescribeSecret(context.Context, *secret.DescribeSecretRequest) (
	*secret.Secret, error) {
	ins := secret.NewDefaultSecret()
	ins.Data.ApiKey = ""
	ins.Data.ApiSecret = ""
	return ins, nil
}

func (m *secretMock) DeleteSecret(context.Context, *secret.DeleteSecretRequest) (*secret.Secret, error) {
	return nil, nil
}
