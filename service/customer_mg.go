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
