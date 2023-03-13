package ecs

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)
import "github.com/playmood/cmdb/apps/host"

func NewPagger(op *EcsOperator) host.Pagger {
	req := &model.ListServersDetailsRequest{}
	p := &pagger{
		op:       op,
		hasNext:  true,
		req:      req,
		pageNum:  1,
		pageSize: 20,
		log:      zap.L().Named("ecs"),
	}
	p.req = req
	p.req.Offset = p.getOffset()
	p.req.Limit = &p.pageSize
	return p
}

type pagger struct {
	req      *model.ListServersDetailsRequest
	op       *EcsOperator
	hasNext  bool
	pageNum  int32
	pageSize int32
	log      logger.Logger
}

func (p *pagger) SetPageSize(ps int32) {
	p.pageSize = ps
}

func (p *pagger) getOffset() *int32 {
	offset := (p.pageNum - 1) * p.pageSize
	return &offset
}

func (p *pagger) Next() bool {
	return p.hasNext
}

func (p *pagger) nextReq() *model.ListServersDetailsRequest {
	p.req.Offset = p.getOffset()
	return p.req
}

func (p *pagger) Scan(ctx context.Context, hs *host.HostSet) error {
	p.log.Debugf("query page: %d", p.pageNum)
	set, err := p.op.QueryInstance(p.nextReq())
	*hs = *set
	if err != nil {
		return err
	}

	// 看当前页是否为满页
	if hs.Length() < int64(p.pageSize) {
		p.hasNext = false
	}

	// 修改指针到下一页
	p.pageNum++

	return nil
}

//type Pagger interface {
//	Next() bool
//	SetPageSize(ps int64)
//	Scan(context.Context, *HostSet) error
//}
