package models

import "gRPC/dao"

// CreateHaving 需要去把菜品和原材料做匹配
// CreateFh 添加菜品中的原材料
// DeleteFh 删除菜品中的原材料
// GetFh 查询菜品原材料映射
// UpdateHaving 修改菜品描述或名字
// DeleteHaving 删除菜品
// GetHavingFood 查询菜品中所有的原材料
// GetHavingIds 按Id查询所有菜品
// GetHaving 按条件检索


// CreateHaving 需要去把菜品和原材料做匹配
func CreateHaving(having *Having,FH []Fh)(err error){
	tx := dao.Db.Begin()
	if err := tx.Create(having).Error;err!=nil{
		tx.Rollback()
		return err
	}
	for i := 0; i < len(FH); i++ {
		FH[i].HavingId = having.HavingId
	}
	if err := tx.Create(FH).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// CreateFh 添加菜品中的原材料
func CreateFh(havingId int,Foods []Fh)error{
	food_data := make([]Fh,0)
	tx := dao.Db
	//tx := dao.Db.Where("_id = ? ",orderId)
	for i:=0 ; i<len(Foods) ;i++ {
		temp := Fh{
			FoodId : Foods[i].FoodId,
			HavingId: havingId,
			FoodWeight: Foods[i].FoodWeight,
		}
		food_data = append(food_data,temp)
	}
	err := tx.Create(food_data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteFh 删除菜品中的原材料
func DeleteFh(havingId int,Foods []Fh)error{
	if len(Foods) == 0{
		return nil
	}
	tx := dao.Db.Where("(having_id = ? and food_id = ? and food_weight = ?)",havingId,Foods[0].FoodId,Foods[0].FoodWeight)
	for i:=1 ; i<len(Foods) ;i++ {
		tx = tx.Or("(having_id = ? and food_id = ? and food_weight = ?)",havingId,Foods[0].FoodId,Foods[0].FoodWeight)
	}
	err := tx.Delete(Fh{}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetFh 查询菜品原材料映射
func GetFh(havingId int)(fh *[]Fh,err error){
	tx := dao.Db.Where("having_id = ? ",havingId)
	errF := tx.Find(fh).Error
	if errF != nil {
		return nil,errF
	}
	return  fh,nil
}

// Findhaving
func FindHaving(limit int,page int)(err error, havings []*Having){
	if err = dao.Db.Limit(limit).Offset((page-1)*limit).Find(&havings).Error;err!=nil{
		return err,nil
	}
	return err,havings
}

// UpdateHaving 修改菜品描述或名字
func UpdateHaving(having *Having)(err error){
	if err = dao.Db.Save(having).Where("having_id = ?",having.HavingId).Error;err != nil{
		return err
	}
	return
}

// DeleteHaving 删除菜品
func DeleteHaving(ID int)(err error){
	if err = dao.Db.Where("having_id = ?",ID).Delete(Having{}).Error;err!=nil{
		return err
	}
	return
}

// GetHaving 按条件检索
func GetHaving(limit int,page int,having Having)(havings *[]Having,err error){
	tx := dao.Db.Where("having_status = ?",1)
	if having.HavingPrice != 0{
		tx = tx.Where("having_price > ?",having.HavingPrice)
	}
	if having.HavingName != ""{
		tx = tx.Where("having_name like %?%",having.HavingName)
	}
	err = tx.Limit(limit).Offset((page-1)*limit).Find(&havings).Error
	if err != nil {
		return nil,err
	}
	return havings,nil
}

// GetHavingIds 按Id查询所有菜品
func GetHavingIds(havingIds []int)(*[]Having,error)  {
	having_data := make([]Having,len(havingIds))
	if len(havingIds) == 0{
		return nil,nil
	}
	tx := dao.Db.Where("having_id = ? ",havingIds[0])
	for i:=0 ; i<len(havingIds) ;i++ {
		tx = tx.Or("having_id = ?",havingIds[i])
	}
	errF := tx.Find(&having_data).Error
	if errF != nil {
		return nil,nil
	}
	return &having_data,nil
}


// GetHavingFood 查询菜品中所有的原材料
func GetHavingFood(HavingId int)(Foods *[]Food,errh error){

	fh , errf := GetFh(HavingId)
	if errf != nil {
		return nil,errf
	}
	if len(*fh) == 0{
		return nil,nil
	}
	tx := dao.Db.Where("food_id",(*fh)[0].FoodId)
	for i := 1; i < len(*fh); i++ {
		tx = dao.Db.Or("food_id",(*fh)[0].FoodId)
	}
	tx.Find(Foods)
	return Foods ,nil
}

// GetHavingOrder 查询订单中所有的菜品
func GetHavingOrder(OrderId int)(Havings *[]Having,errh error){

	oh , errf := GetOh(OrderId)
	if errf != nil {
		return nil,errf
	}
	if len(*oh) == 0{
		return nil,nil
	}
	tx := dao.Db.Where("having_id",(*oh)[0].HavingId)
	for i := 1; i < len(*oh); i++ {
		tx = dao.Db.Or("having_id",(*oh)[0].HavingId)
	}
	tx.Find(Havings)
	return Havings ,nil
}
