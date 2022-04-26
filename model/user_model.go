package models

type User struct {
	UserId int `gorm:"primary_key" json:"user_id"`
	UserName string	`gorm:"type:varchar(255)" json:"user_name"`
	UserPassword string `gorm:"type:varchar(255)" json:"user_password"`
	UserEmail string	`gorm:"type:varchar(255);unique_index" json:"user_email"`
	UserStatus int `gorm:"type:int" json:"user_status"`
	UserPid string `gorm:"type:varchar(255)" json:"user_pid"`
}
func (User) TableName() string {
	return "user"
}

type UserRegister struct {
	UserId int `gorm:"primary_key" json:"user_id"`
	UserName string	`gorm:"type:varchar(255)" json:"user_name"`
	UserPassword string `gorm:"type:varchar(255)" json:"user_password"`
	UserEmail string	`gorm:"type:varchar(255);unique_index" json:"user_email"`
	UserStatus int `gorm:"type:int" json:"user_status"`
	UserSessionId int `json:"user_session_id"`
}