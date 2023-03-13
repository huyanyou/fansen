package controller

import (
	"fansenbk/model/res"
	"fansenbk/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Table struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}

// 创建客户
func CreateCustomer(c *gin.Context) {
	var req service.CreateCustomerReq
	err := c.ShouldBindJSON(&req)
	fmt.Println(req)
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

// 批量删除客户
func DelCustomeres(c *gin.Context) {
	var req service.DelCustomerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: err.Error(),
		})
		return
	}
	cuses, err := req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "删除客户失败",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "删除客户成功",
		Data: cuses,
	})
}

// 批量更新消费者
func UpdateCustomeres(c *gin.Context) {
	var cuses []service.UpCustomeres
	err := c.ShouldBindJSON(&cuses)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: err.Error(),
		})
		return
	}
	req := service.UpdateCustomerReq{
		Customeres: cuses,
	}
	result, err := req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "更新客户失败",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "更新客户成功",
		Data: result,
	})
}
