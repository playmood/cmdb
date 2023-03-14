package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/playmood/cmdb/apps/book/api"
	_ "github.com/playmood/cmdb/apps/host/api"
	_ "github.com/playmood/cmdb/apps/resource/api"
	_ "github.com/playmood/cmdb/apps/secret/api"
	_ "github.com/playmood/cmdb/apps/task/api"
)
