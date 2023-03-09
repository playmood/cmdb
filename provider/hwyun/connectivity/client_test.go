package connectivity_test

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/playmood/cmdb/provider/hwyun/connectivity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHuaweiCloudClient(t *testing.T) {
	should := assert.New(t)
	err := connectivity.LoadClientFromEnv()
	if should.NoError(err) {
		c := connectivity.C()
		fmt.Println(c.Region)
		// 初始化请求,，以调用接口 ListVpcs 为例
		request := &model.ListServersDetailsRequest{}
		client, err := c.EcsClient()
		response, err := client.ListServersDetails(request)
		if err == nil {
			fmt.Printf("%+v\n", response.Servers)
		} else {
			fmt.Println(err)
		}
	}

}
