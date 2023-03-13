package model

type Transaction struct {
	Model
	// 交易类型
	CustomerID uint `json:"customer_id" gorm:"not null foreignkey:ID"`
	// 交易金额
	Amount float64 `json:"amount" gorm:"not null"`
	// 交易类型
	Type uint8 `json:"type" gorm:"not null"`
	// 交易备注
	Remark string `json:"remark"`
	// 交易员
	AdminID string `json:"operator" gorm:"not null"`
}
