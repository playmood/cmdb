package connectivity

import (
	"github.com/caarlos0/env/v6"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	ecs_region "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
)

var (
	client *HuaweiCloudClient
)

func C() *HuaweiCloudClient {
	if client == nil {
		panic("please load config first")
	}
	return client
}

func LoadClientFromEnv() error {
	client = &HuaweiCloudClient{}
	if err := env.Parse(client); err != nil {
		return err
	}

	return nil
}

// NewHuaweiCloudClient client
func NewHuaweiCloudClient(ak, sk, region string) *HuaweiCloudClient {
	return &HuaweiCloudClient{
		Region:       region,
		AccessKey:    ak,
		AccessSecret: sk,
	}
}

type HuaweiCloudClient struct {
	Region       string `env:"HW_CLOUD_REGION"`
	AccessKey    string `env:"HW_CLOUD_ACCESS_KEY"`
	AccessSecret string `env:"HW_CLOUD_ACCESS_SECRET"`

	ecsConn *ecs.EcsClient
}

func (c *HuaweiCloudClient) Credentials() *basic.Credentials {
	auth := basic.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()
	return auth
}

func (c *HuaweiCloudClient) GlobalCredentials() *global.Credentials {
	auth := global.NewCredentialsBuilder().
		WithAk(c.AccessKey).
		WithSk(c.AccessSecret).
		Build()
	return auth
}

// EcsClient 客户端
func (c *HuaweiCloudClient) EcsClient() (*ecs.EcsClient, error) {
	if c.ecsConn != nil {
		return c.ecsConn, nil
	}

	client := ecs.EcsClientBuilder().
		WithRegion(ecs_region.CN_NORTH_4).
		WithCredential(c.Credentials()).
		Build()

	c.ecsConn = ecs.NewEcsClient(client)

	return c.ecsConn, nil
}
