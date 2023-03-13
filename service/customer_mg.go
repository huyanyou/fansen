package service

import (
	"fansenbk/model"

	"github.com/gin-gonic/gin"
)

type CreateCustomerReq struct {
	Role   string  `json:"role" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	Phone  string  `json:"phone" binding:"required"`
	Email  string  `json:"email"`
	Age    int     `json:"age"`
	Remark string  `json:"remark"`
	Remain float64 `json:"remain" binding:"required"`
}

// type UpdateCustomerReq struct {
// 	Customeres []struct {
// 		CreateCustomerReq `json:",inline"`
// 		ID                uint `json:"id" binding:"required"`
// 	} `json:"customeres" binding:"required"`
// }
type UpCustomeres struct {
	CreateCustomerReq
	ID uint `json:"id" binding:"required"`
}
type UpdateCustomerReq struct {
	Customeres []UpCustomeres
}

type DelCustomerReq struct {
	Ids []uint `json:"ids" binding:"required"`
}

// 创建消费者
func (req *CreateCustomerReq) Do(c *gin.Context) (cus model.Customer, err error) {
	cus = model.Customer{
		Role:   req.Role,
		Name:   req.Name,
		Phone:  req.Phone,
		Email:  req.Email,
		Age:    req.Age,
		Remark: req.Remark,
		Remain: req.Remain,
	}
	err = model.CreateCustomer(&cus)
	if err != nil {
		return cus, err
	}
	return cus, err
}

// 批量删除消费者
func (req *DelCustomerReq) Do(c *gin.Context) (cuses []model.Customer, err error) {
	cuses = make([]model.Customer, len(req.Ids))
	err = model.DeleteCustomersByIds(&cuses, req.Ids)
	return cuses, err
}

// 批量更新消费者
func (req *UpdateCustomerReq) Do(c *gin.Context) (cuses []*model.Customer, err error) {
	ids := make([]uint, len(req.Customeres))
	cuses = make([]*model.Customer, len(req.Customeres))
	for i, v := range req.Customeres {
		ids[i] = v.ID
		cuses[i] = &model.Customer{
			Model: model.Model{
				ID: v.ID,
			},
		}
	}
	err = model.UpdateCustomersByIds(cuses, ids)
	return cuses, err
}
