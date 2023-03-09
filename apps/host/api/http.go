package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/playmood/cmdb/apps/host"
)

var (
	h = &handler{}
)

type handler struct {
	service host.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(host.AppName)
	h.service = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return host.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.CreateHost).
		Doc("create a host").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Reads(host.Host{}).
		Writes(response.NewData(host.Host{})))

	ws.Route(ws.GET("/").To(h.QueryHost).
		Doc("get all hosts").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Reads(host.QueryHostRequest{}).
		Writes(response.NewData(host.HostSet{})).
		Returns(200, "OK", host.HostSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeHost).
		Doc("describe an host").
		Param(ws.PathParameter("id", "identifier of the host").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Writes(response.NewData(host.Host{})).
		Returns(200, "OK", response.NewData(host.Host{})).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRESTfulApp(h)
}
