package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	// 商户版本设定
	MerchantVersionCollection = "merchant_version"
)

// MerchantVersion 商户版本定义
// 例如: 免费版、专业版、企业版本
type MerchantVersion struct {
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at" bson:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at" bson:"updated_at"`

	// 状态
	Status bool `json:"status"`

	// 版本名称
	Name string `json:"name"`

	// 小标题
	Subtitle string `json:"sub_title"`

	// 标签
	Badge string `json:"badge"`

	// 应用价格
	Price uint `json:"price" bson:"price"`

	// 版本介绍
	Desc string `json:"desc"`

	// 包含套餐
	Combo []string `json:"combo" bson:"combo"`

}
