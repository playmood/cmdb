package ecs

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/playmood/cmdb/apps/host"
)

// 查询云服务器详情列表
// 参考文档: https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=ECS&api=ListServersDetails
func (o *EcsOperator) QueryInstance(req *model.ListServersDetailsRequest) (*host.HostSet, error) {
	// set := host.NewHostSet()

	resp, err := o.client.ListServersDetails(req)
	if err != nil {
		return nil, err
	}

	o.log.Debugf(resp.String())
	//set.Total = int64(*resp.Count)
	//set.Items = o.transferInstanceSet(resp.Servers).Items

	return nil, nil
}
