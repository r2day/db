package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	MerchantAppsCollection = "merchant_apps"
)

type MerchantApps struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`

	// 状态
	Status bool `json:"status"`

	// 商户名称
	Name string `json:"name"`

	// 应用价格
	Price uint `json:"price" bson:"price"`

	// 应用介绍
	Desc string `json:"desc"`
}
