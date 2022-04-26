package models

type Food struct {
	FoodId int `gorm:"primary_key" json:"food_id"`
	FoodCw float64 `gorm:"type:float" json:"food_cw"`
	FoodOil float64 `gorm:"type:float" json:"food_oil"`
	FoodEnergy float64 `gorm:"type:float" json:"food_energy"`
	FoodProtein float64 `gorm:"type:float" json:"food_protein"`
	FoodName string `gorm:"type:varchar(255)" json:"food_name"`
	FoodType string `gorm:"type:varchar(255)" json:"food_type"`
}
func (Food) TableName() string {
	return "food"
}

type Fh struct {
	Id int `gorm:"primary_key" json:"id"`
	FoodId int `gorm:"type:int" json:"food_id"`
	HavingId int `gorm:"type:int" json:"having_id"`
	FoodWeight float64 `gorm:"type:float" json:"food_weight"`
}
func (Fh) TableName() string {
	return "fh"
}