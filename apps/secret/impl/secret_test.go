package impl_test

import (
	"context"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	_ "github.com/playmood/cmdb/apps/all"
	"github.com/playmood/cmdb/apps/secret"
	"github.com/playmood/cmdb/conf"
	"os"
	"testing"
)

var (
	ins secret.ServiceServer
)

func TestQuerySecret(t *testing.T) {
	ss, err := ins.QuerySecret(context.Background(), secret.NewQuerySecretRequest())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestDescribeSecret(t *testing.T) {
	ss, err := ins.DescribeSecret(context.Background(), secret.NewDescribeSecretRequest("cg7tsdnvhjvj026tt8fg"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestCreateSecret(t *testing.T) {
	req := secret.NewCreateSecretRequest()
	req.Description = "测试用例"
	req.ApiKey = os.Getenv("HW_CLOUD_ACCESS_KEY")
	req.ApiSecret = os.Getenv("HW_CLOUD_ACCESS_SECRET")
	req.AllowRegions = []string{"*"}
	ss, err := ins.CreateSecret(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func init() {
	// 通过环境变量加载测试配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	// 全局日志对象初始化
	zap.DevelopmentSetup()

	// 数据库配置初始化
	conf.LoadConfigFromToml("../../../etc/config.toml")

	// 初始化所有实例
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
	ins = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
}
