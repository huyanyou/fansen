package service

import (
	"fansenbk/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreateAdminReq struct {
	Role  string `json:"role" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Sex   string `json:"sex" binding:"required"`
}

type UpdateAdminReq uint

type GetAdminByIdReq uint

type DelAdminByIdReq uint

func (req *CreateAdminReq) Do(ctx *gin.Context) (admin model.Admin, err error) {
	admin = model.Admin{
		Role:  req.Role,
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
		Age:   req.Age,
	}
	err = model.CreateAdmin(&admin)

	return admin, err
}

// 根据Id获取管理员
func (req *GetAdminByIdReq) Do(ctx *gin.Context) (admin model.Admin, err error) {
	admin = model.Admin{
		Model: model.Model{
			ID: uint(*req),
		},
	}
	err = model.GetAdminById(&admin)
	return admin, err
}

// 根据Id更新管理员
func (req *UpdateAdminReq) Do(ctx *gin.Context) (admin model.Admin, err error) {
	err = ctx.ShouldBindJSON(&admin)
	if err != nil {
		return admin, err
	}
	admin.ID = uint(*req)
	fmt.Println(admin.Age)
	err = model.UpdateAdminById(&admin)
	return admin, err
}

// 根据Id删除管理员
func (req *DelAdminByIdReq) Do(ctx *gin.Context) (admin model.Admin, err error) {
	admin.ID = uint(*req)
	err = model.DeleteAdminById(&admin)
	return admin, err
}
