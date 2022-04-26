package models

type Menu struct {
	MenuId int `gorm:"primary_key" json:"menu_id"`
	CreateTime string `gorm:"type:datetime" json:"create_time"`
	MenuStatus string `gorm:"type:varchar(255)" json:"menu_status"`
}
func (Menu) TableName() string {
	return "menu"
}
type Mh struct {
	Id int `gorm:"primary_key" json:"id"`
	MenuId int `gorm:"type:int" json:"menu_id"`
	HavingId int `gorm:"type:int" json:"having_id"`
}
func (Mh) TableName() string {
	return "mh"
}