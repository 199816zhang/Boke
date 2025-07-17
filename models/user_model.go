package models

type UserModel struct {
	Model
	Username       string `gorm:"size:32" json:"username"`
	Nickname       string `gorm:"size:32" json:"nickname"`
	Avatar         string `gorm:"size:256" json:"avatar"`   //头像
	Abstract       string `gorm:"size:256" json:"abstract"` //简介
	RegisterSource int8   `json:"registerSource"`           // 注册来源
	CodeAge        int    `json:"codeAge"`                  // 码龄
	Password       string `gorm:"size:64" json:"-"`
	Email          string `gorm:"size:256" json:"email"`
	OpenID         string `gorm:"size:64" json:"openID"` // 第三方登陆的唯一id
}
