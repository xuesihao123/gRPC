package models

import (
	"gRPC/dao"
	"time"
)

const (
	DayDate = "DayDate"
	OftenDate = "OftenDate"
)

// CreateMenu 创建今日菜单
// CreateMenuHaving 菜单新增菜品
// DeleteMh 菜单删除菜品
// GetMh 查询菜单菜品
// GetMenuHavingDate 按日期查询菜单中所有的菜品


// CreateMenu 创建今日菜单
func CreateMenu(havingIds []int32)(err error){
	_ , errG := GetMenuDate(time.Now().Format("2006-01-02"))
	if errG == nil {
		return nil  //应该返回已有
	}
	tx := dao.Db.Begin()
	menu := &Menu{
		CreateTime: time.Now().Format("2006-01-02"),
		MenuStatus: DayDate,
	}
	errC := tx.Create(menu).Error
	if errC != nil {
		tx.Rollback()
		return errC
	}
	menu_id := menu.MenuId

	order_data := make([]Mh,0)

	for i := 0; i < len(havingIds); i++ {
		temp := Mh{
			MenuId : menu_id,
			HavingId: int(havingIds[i]),
		}
		order_data = append(order_data,temp)
	}
	if errD := tx.Create(order_data).Error;err!=nil{
		tx.Rollback()
		return errD
	}
	tx.Commit()
	return nil
}

// CreateMenuHaving 菜单新增菜品
func CreateMenuHaving(menu Menu,havings []int)error{
	order_data := make([]Mh,0)
	//tx := dao.Db.Where("menu_id = ? ",menu.MenuId)
	tx := dao.Db
	for i:=0 ; i<len(havings) ;i++ {
		temp := Mh{
			MenuId : menu.MenuId,
			HavingId: havings[i],
		}
		order_data = append(order_data,temp)
	}
	err := tx.Create(order_data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteMh 菜单删除菜品
func DeleteMh(menuId int,havings []int32)error{
	if len(havings) == 0{
		return nil
	}
	tx := dao.Db.Where("(menu_id = ? and having_id = ?)",menuId,havings[0])
	for i:=1 ; i<len(havings) ;i++ {
		tx = tx.Or("(menu_id = ? and having_id = ?)",menuId,havings[i])
	}
	err := tx.Delete(Oh{}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetMh 查询菜单菜品
func GetMh(menuId int)(mh *[]Mh,err error){
	tx := dao.Db.Where("menu_id = ? ",menuId)
	errF := tx.Find(mh).Error
	if errF != nil {
		return nil,errF
	}
	return  mh,nil
}


// GetMenuDate 按照日期查询
func GetMenuDate(Date string)(menu *Menu,err error)  {
	menu = new(Menu)
	if errC := dao.Db.Debug().Where("create_time = ? ",Date).First(menu).Error;errC!=nil{
		return nil,errC
	}
	return menu,nil
}

// GetMenuHavingDate 按日期查询菜单中所有的菜品
func GetMenuHavingDate(Date string)(havings *[]Having,errh error){
	menu , err := GetMenuDate(Date)
	if err != nil {
		return nil,err
	}
	mh , errm := GetMh(menu.MenuId)
	if errm != nil {
		return nil,err
	}
	if len(*mh) == 0{
		return nil,nil
	}
	tx := dao.Db.Where("having_id",(*mh)[0].HavingId)
	for i := 1; i < len(*mh); i++ {
		tx = dao.Db.Or("having_id",(*mh)[0].HavingId)
	}
	tx.Find(havings)
	return havings ,nil
}

// CreateOftenMenu 创建常用菜单
func CreateOftenMenu(havingIds []int32)(err error){
	menu_id := 1
	order_data := make([]Mh,0)

	for i := 0; i < len(havingIds); i++ {
		temp := Mh{
			MenuId : menu_id,
			HavingId: int(havingIds[i]),
		}
		order_data = append(order_data,temp)
	}
	if errD := dao.Db.Create(order_data).Error;err!=nil{
		return errD
	}
	return nil
}

// DeleteOftenMh 菜单删除菜品
func DeleteOftenMh(menuId int,havings []int32)error{
	menuId = 1
	if len(havings) == 0{
		return nil
	}
	tx := dao.Db.Where("(menu_id = ? and having_id = ?)",menuId,havings[0])
	for i:=1 ; i<len(havings) ;i++ {
		tx = tx.Or("(menu_id = ? and having_id = ?)",menuId,havings[i])
	}
	err := tx.Delete(Oh{}).Error
	if err != nil {
		return err
	}
	return nil
}
//菜单的周期性和时效性
//菜品的营养成分确认
//新增对于原有食材的营养成分维护