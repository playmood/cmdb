package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/playmood/cmdb/apps/task"
)

func (h *handler) CreatTask(r *restful.Request, w *restful.Response) {
	req := task.NewCreateTaskRequst()
	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	r.Request.BasicAuth()
	// 直接启动一个goroutine 来执行,
	// 想要通过Task做异常, 这里需要改造, 支持传递Task Id 参数
	// go func() {
	// 	set, err := h.task.CreateTask(r.Request.Context(), req)
	// }()

	set, err := h.task.CreateTask(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryTask(r *restful.Request, w *restful.Response) {
	// query := task.NewQueryTaskRequestFromHTTP(r.Request)
	// set, err := h.task.QueryTask(r.Request.Context(), query)
	// if err != nil {
	// 	response.Failed(w, err)
	// 	return
	// }

	response.Success(w, nil)
}

func (h *handler) DescribeTask(r *restful.Request, w *restful.Response) {
	// req := task.NewDescribeTaskRequestWithId(r.PathParameter("id"))
	// ins, err := h.task.DescribeTask(r.Request.Context(), req)
	// if err != nil {
	// 	response.Failed(w, err)
	// 	return
	// }

	response.Success(w, nil)
}
