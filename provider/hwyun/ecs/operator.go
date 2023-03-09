package ecs

import (
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewEcsOperator(client *ecs.EcsClient) *EcsOperator {
	return &EcsOperator{
		client: client,
		log:    zap.L().Named("hw.ecs"),
	}
}

type EcsOperator struct {
	client *ecs.EcsClient
	log    logger.Logger
}
