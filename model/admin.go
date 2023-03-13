package model

import (
	"errors"

	"gorm.io/gorm/clause"

	"fansenbk/utils"
)

type Admin struct {
	Model
	// 角色
	Role string `json:"role"`
	// 姓名
	Name string `json:"name"`
	// 性别
	Sex string `json:"sex"`
	// 年龄
	Age int `json:"age"`
	// 电话
	Phone string `json:"phone" gorm:"not null;unique"`
	// 邮箱
	Email string `json:"email" gorm:"not null;unique"`
}

// 创建管理员
func CreateAdmin(ad *Admin) (err error) {
	// 创建管理员,如果创建失败,则返回错误
	err = DB.Table("admins").Create(&ad).Error
	if err != nil {
		// 判断是否是唯一性错误
		return errors.New(utils.IsUniqueConstraintError(err))
	}
	// 获取创建的管理员
	return err
}

// 根据Id获取管理员
func GetAdminById(ad *Admin) (err error) {
	err = DB.Table("admins").First(&ad, ad.ID).Error
	return err
}

// 删除管理员
func DeleteAdminById(ad *Admin) (err error) {
	result := DB.Clauses(clause.Returning{}).Table("admins").Delete(&ad)
	if result.RowsAffected == 0 {
		return errors.New("管理员不存在")
	}
	return result.Error
}

// 更新管理员
func UpdateAdminById(ad *Admin) (err error) {
	// 查看管理员是否成存在
	result := DB.Model(&Admin{}).Where("id = ?", ad.ID).First(&ad)
	if result.RowsAffected == 0 {
		return errors.New("管理员不存在")
	}
	DB.Model(ad).Updates(Admin{
		Name:  ad.Name,
		Email: ad.Email,
		Sex:   ad.Sex,
		Age:   ad.Age,
		Phone: ad.Phone,
		Role:  ad.Role,
	})
	return err
}

// 获取管理员列表
func GetAdminList(ads *[]Admin) (err error) {
	err = DB.Table("admins").Find(&ads).Error
	return err
}
