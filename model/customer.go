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
		if utils.IsUniqueCustomerError(err) != "" {
			return errors.New(utils.IsUniqueCustomerError(err))
		}
		return err
	}
	return err
}

// 批量删除客户
func DeleteCustomersByIds(cus *[]Customer, ids []uint) (err error) {
	// 查询总共有多少条记录
	// DB.Table("customers").Count(&total)
	result := DB.Table("customers").Find(&[]Customer{}, ids)
	// 查询ids所有用户
	if int(result.RowsAffected) != len(*cus) {
		return errors.New("客户不存在")
	}
	result = DB.Clauses(clause.Returning{}).Table("customers").Delete(&cus, ids)
	return result.Error
}

// 批量更新客户
func UpdateCustomersByIds(cuses []*Customer, ids []uint) (err error) {
	// 判断用户是否存在
	result := DB.Find(&[]Customer{}, ids)
	if result.Error != nil {
		return result.Error
	}
	if int(result.RowsAffected) != len(cuses) {
		return errors.New("客户不存在")
	}
	DB.Clauses(clause.Returning{}).Table("customers").Updates(&cuses)
	return result.Error
}
