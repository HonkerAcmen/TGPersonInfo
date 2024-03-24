package model

type UserInfo struct {
	ID        int64  `gorm:"primarykey"`
	UserID    string `gorm:"column:userid" json:"UserID"`
	UserName  string `gorm:"column:username" json:"UserName"`
	UserSign  string `gorm:"column:usersign" json:"UserSign"`
	Following int64  `gorm:"column:following" json:"Following"`
	Followers int64  `gorm:"column:followers" json:"Followers"`
}
