package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountManager struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 商户号
	MerchantId string `json:"-" bson:"merchant_id"`
	// 门店号
	StoreId string `json:"-" bson:"store_id"`
	// 创建者
	UserId string `json:"user_id" bson:"user_id"`
	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`
	// 更新人
	UpdatedUserId string `json:"updated_user_id" bson:"updated_user_id"`
	// 状态
	Status bool `json:"status"`
	// 账号名称 admin
	Name string `json:"name"`
	// 账号Id admin
	AccountId string `json:"account_id"  bson:"account_id"`

	// 手机号
	Phone string `json:"phone"`
	// 角色类型
	// 根据角色列表
	// 合并权限列表（apps)
	// 例如: 店长/后厨 （可以查看两者的所有数据）
	// 保存的是角色的id
	// 角色访问app 变后，这里会自动关联
	Roles []string `json:"roles"  bson:"roles,omitempty"`
	// 角色描述
	Desc string `json:"desc"`
	// 角色来源
	RoleFrom string `json:"role_from" bson:"role_from,omitempty"`

	// 账号分组
	GroupId string `json:"group_id" bson:"group_id,omitempty"`

}
