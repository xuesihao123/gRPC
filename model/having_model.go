package models

type Having struct {
	HavingId int `gorm:"primary_key" json:"having_id""`
	HavingPrice float64 `gorm:"type:float" json:"having_price"`
	HavingName string `gorm:"type:varchar(25)" json:"having_name"`
	HavingCW float64 `gorm:"type:float" json:"having_cw"`
	HavingOil float64 `gorm:"type:float" json:"having_oil"`
	HavingEnergy float64 `gorm:"type:float" json:"having_energy"`
	HavingProtein float64 `gorm:"type:float" json:"having_protein"`
	HavingUpdateTime string `gorm:"type:datetime" json:"having_update_time"`
	HavingContent string `gorm:"type:varchar(255)" json:"having_content"`
}
func (Having) TableName() string {
	return "having"
}