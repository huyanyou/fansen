package controller

import (
	"fansenbk/model/res"
	"fansenbk/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建客户
func CreateCustomer(c *gin.Context) {
	var req service.CreateCustomerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: err.Error(),
		})
		return
	}
	cus, err := req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "创建客户失败",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "创建客户成功",
		Data: cus,
	})
}
