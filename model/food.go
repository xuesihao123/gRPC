package models

import "gRPC/dao"

//新增食材
//修改食材
//删除食材
//按名字查询食材

// CreateFood 新增食材
func CreateFood(food *[]Food)(err error){
	if err := dao.Db.Create(food).Error;err!=nil{
		return err
	}
	return nil
}

// UpdateFood 修改食材
func UpdateFood(food *Food)(err error){
	if err = dao.Db.Save(food).Where("food_id = ?",food.FoodId).Error;err != nil{
		return err
	}
	return
}

// DeleteFood 删除食材
func DeleteFood(ID int)(err error){
	if err = dao.Db.Where("food_id = ?",ID).Delete(Food{}).Error;err!=nil{
		return err
	}
	return
}

// GetFood 按名字和种类查询食材
func GetFood(food Food,limit int,page int)(foods *[]Food,err error){
	tx := dao.Db
	if food.FoodType != ""{
		tx = tx.Where("having_type = ?",food.FoodType)
	}
	if food.FoodName != ""{
		tx = tx.Where("having_name like %?%",food.FoodName)
	}
	err = tx.Limit(limit).Offset((page-1)*limit).Find(&foods).Error
	if err != nil {
		return nil,err
	}
	return foods,nil
}

