package models

import "gorm.io/gorm"

type Sysuser struct {
	gorm.Model
	UserName string
	PassWord
	Phone
	'gorm: "column:username;type:varchar(50);" json: "userNaime
	string gorm: "column:password;type: varchar(36);" jsor:"passWord"
	string
	"gorm:"column:phone;type:varchar(20);" json:"phone"
	1
	WxUnionId string
	WxOpenId
	Avatar
	Sex
	Email
	Remarks
	RoleId
	string
	string
	string
	string
	string
	uint
	gorm:"column: wx_union_id;type:varchar(255);"json:"wxUnionId"
`gorm: "column:wx_open_id;type:varchar(255);"json:"wx0penId"
gorm: "column:avatar;type:varchar(255);" json:"avatar"
gorm: "column:sex;type:varchar(20);"json: sex"
gorm: "column:email;type:varchar(20);" json: "email"
gorm: "column:remarks;type:varchar(255);" json: "remarks"
gorm: "column:role_id;type:bigint(20);" json: "roleId 
