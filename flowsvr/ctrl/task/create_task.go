package task

import (
	"github.com/gin-gonic/gin"
	"github.com/wureny/TaskWave/flowsvr/constant"
	"github.com/wureny/TaskWave/taskutils/rpc/model"
	"net/http"
)

type CreateTaskHandler struct {
	Req    model.CreateTaskReq
	Resq   model.CreateTaskResp
	UserId string
}

func (c CreateTaskHandler) HandleInput() error {
	//TODO implement me
	panic("implement me")
}

func (c CreateTaskHandler) HandleProcess() error {
	//TODO implement me
	panic("implement me")
}

func CreateTask(c *gin.Context) {
	var cth CreateTaskHandler
	defer func() {
		cth.Resq.Msg = constant.GetErrMsg(cth.Resq.Code)
		c.JSON(http.StatusOK, cth.Resq)
	}()
	cth.UserId = c.Request.Header.Get(constant.HEADER_USERID)
	if err := c.ShouldBind(&cth.Req); err != nil {
		//TODO: log error
		cth.Resq.Code = constant.ERR_SHOULD_BIND
		return
	}
	err := cth.HandleInput()
	if err != nil {
		//TODO: log error
		return
	}
	err = cth.HandleProcess()
	if err != nil {
		//TODO: log error
	}
}
