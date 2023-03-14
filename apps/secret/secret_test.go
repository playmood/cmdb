package secret_test

import (
	"testing"

	"github.com/playmood/cmdb/apps/secret"
)

var (
	encryptKey = "asdfgqwer"
)

func TestSecretEncrypt(t *testing.T) {
	ins := secret.NewDefaultSecret()
	ins.Data.ApiSecret = "234567890"
	ins.Data.EncryptAPISecret(encryptKey)
	t.Log(ins.Data.ApiSecret)

	ins.Data.DecryptAPISecret(encryptKey)
	t.Log(ins.Data.ApiSecret)
}
