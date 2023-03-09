package ecs_test

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/playmood/cmdb/provider/hwyun/connectivity"
	"github.com/playmood/cmdb/provider/hwyun/ecs"
	"testing"
)

var (
	op *ecs.EcsOperator
)

func TestQuery(t *testing.T) {
	req := &model.ListServersDetailsRequest{}
	set, err := op.QueryInstance(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	err := connectivity.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	zap.DevelopmentSetup()
	client, err := connectivity.C().EcsClient()
	if err != nil {
		panic(err)
	}
	op = ecs.NewEcsOperator(client)
}
