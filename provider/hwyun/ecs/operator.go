package ecs

import (
	ecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/playmood/cmdb/apps/host"
	"github.com/playmood/cmdb/apps/resource"
	"github.com/playmood/cmdb/utils"
	"strconv"
	"strings"
	"time"
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

func (o *EcsOperator) transferInstanceSet(list *[]model.ServerDetail) *host.HostSet {
	set := host.NewHostSet()
	items := *list
	for i := range items {
		set.Add(o.transferInstance(items[i]))
	}
	return set
}

func (o *EcsOperator) transferInstance(ins model.ServerDetail) *host.Host {
	r := host.NewDefaultHost()
	b := r.Base
	b.CreateAt = o.parseTime(ins.Created)
	b.Id = ins.Id
	b.Vendor = resource.Vendor_HUAWEI
	b.Zone = ins.OSEXTAZavailabilityZone
	i := r.Information
	i.Category = ins.Flavor.Name
	i.ExpireAt = o.parseTime(ins.AutoTerminateTime)
	i.Name = ins.Name
	i.Description = utils.PtrStrV(ins.Description)
	d := r.Describe
	d.SerialNumber = ins.Id
	cpu, _ := strconv.ParseInt(ins.Flavor.Vcpus, 10, 64)
	d.Cpu = cpu
	mem, _ := strconv.ParseInt(ins.Flavor.Ram, 10, 64)
	d.Memory = mem

	if ins.Tags != nil {
		i.Tags = o.transferTags(*ins.Tags)
	}

	r.Describe.OsType = ins.Metadata["os_type"]
	r.Describe.OsName = ins.Metadata["image_name"]
	r.Describe.ImageId = ins.Image.Id
	r.Describe.KeyPairName = []string{ins.KeyName}
	return r
}

func (o *EcsOperator) parseTime(t string) int64 {
	if t == "" {
		return 0
	}

	ts, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		o.log.Errorf("parse time %s error, %s", t, err)
		return 0
	}

	return ts.UnixNano() / 1000000
}

func (o *EcsOperator) parseIp(address map[string][]model.ServerAddress) (privateIps []string, publicIps []string) {
	for _, addrs := range address {
		for i := range addrs {
			switch *addrs[i].OSEXTIPStype {
			case model.GetServerAddressOSEXTIPStypeEnum().FIXED:
				privateIps = append(privateIps, addrs[i].Addr)
			case model.GetServerAddressOSEXTIPStypeEnum().FLOATING:
				publicIps = append(publicIps, addrs[i].Addr)
			}
		}
	}
	return
}

func (o *EcsOperator) transferTags(tags []string) (ret []*resource.Tag) {
	for _, t := range tags {
		kv := strings.Split(t, "=")
		if len(kv) == 2 {
			ret = append(ret, resource.NewThirdTag(kv[0], kv[1]))
		} else {
			ret = append(ret, resource.NewThirdTag("ecs", t))
		}
	}

	return
}
