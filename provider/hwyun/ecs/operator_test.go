package ecs_test

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/playmood/cmdb/apps/host"
	"github.com/playmood/cmdb/provider/hwyun/connectivity"
	"github.com/playmood/cmdb/provider/hwyun/ecs"
	"testing"
)

var (
	op *ecs.EcsOperator
)

func TestQuery(t *testing.T) {
	var (
		offset int32 = 0
		limit  int32 = 20
	)
	req := &model.ListServersDetailsRequest{}
	req.Offset = &offset
	req.Limit = &limit
	set, err := op.QueryInstance(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestPaggerQuery(t *testing.T) {
	p := ecs.NewPagger(op)
	set := host.NewHostSet()
	for p.Next() {
		if err := p.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		t.Log(set)
	}
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
