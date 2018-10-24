package controllers

import (
	"github.com/gin-gonic/gin"
	"autumn/tools/cfg"
	"strconv"
	"autumn/models"
	"autumn/tools/exception"
	"autumn/common/response"
	"autumn/result"
)

type UserMsgController struct {
	BaseController
}

func (i *UserMsgController) List(c *gin.Context) {
	defer exception.Catch(c)

	page,_ := strconv.Atoi(c.Param("page"))
	limit := int(cfg.Get("env", "page_size").Int())

	offset := (page - 1) * limit
	uid := i.GetUid(c)

	list := (&models.Msg{}).List(i.GetUid(c), limit, offset)

	var sys_msg_id []int
	for _, v := range list {
		sys_msg_id = append(sys_msg_id, v.Id)
	}

	sys_status := (&models.UserSysMsg{}).SysReadDict(uid, sys_msg_id)

	ret := make([]result.MsgList, len(list))
	for k, v := range list {
		status := v.Status
		if _,ok := sys_status[v.Id]; ok {
			status = 1
		}

		item := result.MsgList{v.Id,
			v.Title,
			v.Content,
			status,
			v.CreatedAt,}

		ret[k] = item

	}

	response.Success(c, ret)
}

func (i *UserMsgController) Read(c *gin.Context) {
	defer exception.Catch(c)

	id,_ := strconv.Atoi(c.Param("msg_id"))

	uid := i.GetUid(c)

	if id <= 0 {
		response.Fail(c, 15000)
		return
	}

	info := (&models.Msg{}).Info(id)

	if info.Id == 0 {
		response.Fail(c, 15001)
		return
	}

	if info.Uid == 0 {
		umsg := models.UserSysMsg{}

		//已读
		if umsg.IsRead(uid, info.Id) {
			response.Fail(c, 15002)
			return
		}

		umsg.Uid = uid
		umsg.MsgId = info.Id

		if umsg.Create() {
			response.Success(c, nil)
			return
		}

		response.Fail(c, 1)
		return
	}

	//已读
	if info.Status == 1 {
		response.Fail(c, 15002)
		return
	}

	info.Status = 1

	if info.Update() {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1)
}