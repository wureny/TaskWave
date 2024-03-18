package task

import (
	"github.com/gin-gonic/gin"
	"github.com/wureny/TaskWave/flowsvr/constant"
	"github.com/wureny/TaskWave/flowsvr/db"
	"github.com/wureny/TaskWave/taskutils/rpc/model"
	"net/http"
)

type SetTaskHandler struct {
	Req    model.SetTaskReq
	Resp   model.SetTaskResp
	UserId string
}

func (p SetTaskHandler) HandleInput() error {
	if p.Req.TaskData.TaskId == "" {
		//TODO: log error
		p.Resp.Code = constant.ERR_INPUT_INVALID
		return constant.ERR_HANDLE_INPUT
	}
	if p.Req.TaskData.Priority != nil {
		if *p.Req.TaskData.Priority > db.MAX_PRIORITY || *p.Req.TaskData.Priority < 0 {
			p.Resp.Code = constant.ERR_INPUT_INVALID
			return constant.ERR_HANDLE_INPUT
		}
	}
	return nil
}

func (s SetTaskHandler) HandleProcess() error {
	//TODO implement me
	panic("implement me")
}

func SetTask(c *gin.Context) {
	var hd SetTaskHandler
	defer func() {
		hd.Resp.Msg = constant.GetErrMsg(hd.Resp.Code)
		c.JSON(http.StatusOK, hd.Resp)
	}()
	hd.UserId = c.Request.Header.Get(constant.HEADER_USERID)
	if err := c.ShouldBind(&hd.Req); err != nil {
		//TODO: log error
		hd.Resp.Code = constant.ERR_SHOULD_BIND
		return
	}
	err := hd.HandleInput()
	if err != nil {
		//TODO: log error
		return
	}
	err = hd.HandleProcess()
	if err != nil {
		//TODO: log error
		return
	}
}
