package models

type SysUser struct {
	Usercode string `json:"usercode" gorm:"size:128;comment:编码"`
	Username string `json:"username" gorm:"size:64;comment:用户名"`
	Password string `json:"password" gorm:"size:128;comment:密码"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
