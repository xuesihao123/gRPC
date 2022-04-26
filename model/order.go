package models

import (
	"fmt"
	"gRPC/dao"
	"time"
)

const (
	CreateStatus = "create"
	SuccessStatus = "Success"
	DeleteStatus = "delete"
)

// CreateOrder 新增订单
func CreateOrder(order *OrderJson)(err error){
	tx := dao.Db.Begin()
	OrderOne := &Order{
		UserId:          order.UserId,
		OrderTag:        order.OrderTag,
		OrderPrice:      order.OrderPrice,
		OrderCreateTime: time.Now().Format("2006-01-02 15:04:05"),
		OrderUpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		OrderStatus:     CreateStatus,
		OrderRemark:     "",
	}
	fmt.Printf("%v",OrderOne)
	errC := tx.Create(OrderOne).Error
	if errC != nil {
		tx.Rollback()
		return errC
	}
	order_id := OrderOne.OrderId
	fmt.Printf("%v",OrderOne)
	order_data := make([]Oh,0)

	for i := 0; i < len(order.OrderHavingId); i++ {
		temp := Oh{
			OrderId : order_id,
			HavingId: order.OrderHavingId[i],
		}
		order_data = append(order_data,temp)
	}
	fmt.Printf("%v,%d",order_data,len(order_data))
	//temp := []Oh{{
	//	OrderId : order_id,
	//	HavingId: 1},
	//}
	//errD := tx.Create(&order_data).Error
	errD := tx.Create(&order_data).Error
	//fmt.Printf("122")
	if errD != nil{
		//fmt.Printf("122")
		tx.Rollback()
		return errD
	}
	tx.Commit()
	//fmt.Printf("122")

	return nil
}

// DeleteOrder 删除订单
func DeleteOrder(ID []int)(err error){
	tx := dao.Db.Where("order_id = ? ",ID[0])
	for i:=1 ; i<len(ID) ;i++ {
		tx = tx.Or("order_id = ?",ID[i])
	}
	if err = tx.Save(Order{
		OrderStatus:     DeleteStatus,
	}).Error;err!=nil{
		return err
	}
	return
}

// DeleteOh 删除订单菜品映射
func DeleteOh(orderId int,havings []int32)error{
	if len(havings) == 0{
		return nil
	}
	tx := dao.Db.Where("(order_id = ? and having_id = ?)",orderId,havings[0])
	for i:=1 ; i<len(havings) ;i++ {
		tx = tx.Or("(order_id = ? and having_id = ?)",orderId,havings[i])
	}
	err := tx.Delete(Oh{}).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateOh 创建订单菜品映射
func CreateOh(orderId int,havings []int32)error{

	order_data := make([]Oh,0)
	tx := dao.Db
	for i:=0 ; i<len(havings) ;i++ {
		temp := Oh{
			OrderId : orderId,
			HavingId: int(havings[i]),
		}
		order_data = append(order_data,temp)
	}
	err := tx.Create(order_data).Error
	if err != nil {
		return err
	}
	return nil
}


// GetOh 查询订单菜品映射
func GetOh(orderId int)(oh *[]Oh,err error){
	tx := dao.Db.Where("order_id = ? ",orderId)
	errF := tx.Find(oh).Error
	if errF != nil {
		return nil,errF
	}
	return  oh,nil
}

// GetOrderOfId 查询订单所有的菜品
func GetOrderOfId(orderId int)(Havings *[]Having,err error){
	oh , errf := GetOh(orderId)
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

// GetOrderUserId 查询个人订单
func GetOrderUserId(userId int,limit int,page int)(order *[] Order,err error){
	//var u *User
	if err := dao.Db.Debug().Limit(limit).Offset((page-1)*limit).Where("user_id = ?",userId).Find(order).Error;err!=nil{
		return nil,err
	}
	return order,nil
}

// GetStatusCreate 查询未完成订单
func GetStatusCreate(limit int,page int)(order *[] Order,err error){
	//var u *User
	if err := dao.Db.Debug().Limit(limit).Offset((page-1)*limit).Where("order_status = ?",CreateStatus).Find(order).Error;err!=nil{
		return nil,err
	}
	return order,nil
}

// UpdateOrder 修改订单
func UpdateOrder(order *Order)(err error){
	if err = dao.Db.Save(order).Where("order_id = ?",order.OrderId).Error;err != nil{
		return err
	}
	return
}

// GetOrder 按条件查询订单
func GetOrder(limit int,page int,order *Order)(Orders []*Order ,err error){
	tx := dao.Db
	if order.UserId != 0{
		tx = tx.Where("user_id = ?",order.UserId)
	}
	if order.OrderId != 0{
		tx = tx.Where("order_id = ?",order.UserId)
	}
	if order.OrderStatus != ""{
		tx = tx.Where("order_Status = ?",order.OrderStatus)
	}
	err = tx.Debug().Find(&Orders).Error
	//fmt.Printf("%v",Orders)
	if err != nil {
		return nil,nil
	}
	return Orders, err
}