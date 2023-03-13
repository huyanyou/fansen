package model

import (
	"errors"
	"fansenbk/utils"

	"gorm.io/gorm/clause"
)

type Customer struct {
	Model
	// 姓名
	Name  string `json:"name" gorm:"not null"`
	Phone string `json:"phone" gorm:"not null;unique"`
	// 性别
	Sex string `json:"sex" gorm:"default:'男'"`
	// 余额
	Remain float64 `json:"remain" gorm:"default:0"`
	// 备注
	Remark string `json:"remark"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Role   string `json:"role"`
}

// 创建客户
func CreateCustomer(cus *Customer) (err error) {
	// 创建客户,如果创建失败,则返回错误
	err = DB.Table("customers").Create(&cus).Error
	if err != nil {
		// 判断是否是唯一性错误
		return errors.New(utils.IsUniqueConstraintError(err))
	}
	// 获取创建的客户
	return err
}

// 批量删除客户
func DeleteCustomersByIds(cus *Customer, ids []uint) (err error) {
	result := DB.Clauses(clause.Returning{}).Table("customers").Delete(&cus, ids)
	if result.RowsAffected == 0 {
		return errors.New("客户不存在")
	}
	return result.Error
}
