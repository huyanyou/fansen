package controller

import (
	"fansenbk/model/res"
	"fansenbk/service"
	"fansenbk/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 创建管理员用户
func CreateAdmin(c *gin.Context) {
	var admin_req service.CreateAdminReq
	err := c.ShouldBindJSON(&admin_req)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: err.Error(),
		})
		return
	}
	admin, err := admin_req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   utils.IsUniqueConstraintError(err),
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "创建管理员成功",
		Data: admin,
	})
}

// 根据Id更新管理员用户
func UpdateAdminById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: "",
		})
		return
	}
	// string 转换成 uint
	id_uint, _ := strconv.Atoi(id)
	req := service.UpdateAdminReq(id_uint)
	admin, err := req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   utils.IsUniqueConstraintError(err),
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "更新管理员成功",
		Data: admin,
	})
}

// 根据id获取管理员用户
func GetAdminById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: "",
		})
		return
	}
	id_uint, _ := strconv.Atoi(id)
	req := service.GetAdminByIdReq(id_uint)
	admin, err := req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "管理员不存在",
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "获取管理员成功",
		Data: admin,
	})
}

// 删除管理员用户
func DeleteAdminById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Msg:   "参数错误",
			Error: "",
		})
		return
	}
	id_uint, _ := strconv.Atoi(id)
	req := service.DelAdminByIdReq(id_uint)
	admin, err := req.Do(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		Code: http.StatusOK,
		Msg:  "删除管理员成功",
		Data: admin,
	})
}
