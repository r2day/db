package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	RoleManagerCollection = "role_manage"
)


type RoleManager struct {
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
	// 角色名称
	Name string `json:"name"`
	// 角色类型
	Roles []string `json:"roles"  bson:"roles,omitempty"`
	// 授权的应用
	Apps []string `json:"apps"  bson:"apps,omitempty"`
	// 角色描述
	Desc string `json:"desc"`
	// 角色来源
	RoleFrom string `json:"role_from" bson:"role_from,omitempty"`

}
