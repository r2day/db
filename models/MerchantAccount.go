package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	MerchantAccountCollection = "merchant_account"
)

type MerchantAccount struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`

	// 状态
	Status bool `json:"status"`

	// 商户ID
	MerchantId string `json:"merchant_id" bson:"merchant_id"`

	// 商户名称
	Name string `json:"name"`

	// 版本
	// 例如: 试用、专业、企业
	// 保存版本id
	Version string `json:"version"`

	// 商户手机号
	Phone string `json:"phone"`

	// 最大允许开设的店铺数量
	MaxStoreNumber uint `json:"max_store_number" bson:"max_store_number"`

	// 当前开设的店铺数量
	StoreNumber uint `json:"store_number" bson:"store_number"`

	// 密码
	Password string `json:"password"`

	// 商户介绍
	Desc string `json:"desc"`

	// 品牌Logo
	LogoUrl string `json:"logo_url" bson:"logo_url,omitempty"`

}
