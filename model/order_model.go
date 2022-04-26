package models

type Order struct {
	OrderId int `gorm:"primary_key" json:"order_id""`
	UserId int `gorm:"type:int" json:"user_id""`
	OrderTag string `gorm:"type:varchar(255)" json:"order_tag"`
	OrderPrice float64 `gorm:"type:float" json:"order_price"`
	OrderCreateTime string `gorm:"type:datetime" json:"order_create_time"`
	OrderUpdateTime string `gorm:"type:datetime" json:"order_update_time"`
	OrderStatus string `gorm:"type:varchar(255)" json:"order_status"`
	OrderRemark string `gorm:"type:varchar(255)" json:"order_remark"`
}
func (Order) TableName() string {
	return "order"
}
type OrderJson struct {
	OrderId int `gorm:"primary_key" json:"order_id""`
	UserId int `gorm:"type:int" json:"user_id""`
	OrderTag string `gorm:"type:varchar(255)" json:"order_tag"`
	OrderPrice float64 `gorm:"type:float" json:"order_price"`
	OrderHavingId []int `gorm:"type:int" json:"order_having_id"`
}

type Oh struct {
	Id int `gorm:"primary_key" json:"id""`
	OrderId int `gorm:"type:int" json:"order_id"`
	HavingId int `gorm:"type:int" json:"having_id"`
}
func (Oh) TableName() string {
	return "oh"
}