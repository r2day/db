package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FunctionRequest Binding from JSON
type FunctionRequest struct {
	// 商户号
	MerchantId string `form:"merchant_id" json:"merchant_id" xml:"merchant_id"  binding:"required"`
}


type FunctionDetail struct {
	// 状态
	Status bool `json:"status"`

	// 功能名称
	Name string `json:"name"`
}

// MerchantAppConf 
type MerchantAppConf struct {
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

	// 功能名称
	Name string `json:"name"`

	MenuEnableConf MenuEnable `json:"menu_enable_conf" bson:"menu_enable_conf"`

	MenuDisplayConf MenuDisplay  `json:"menu_display_conf" bson:"menu_display_conf"`
}

type MenuEnable struct {
	// 是否开通/启用 dash
	EnableDash bool `json:"enable_dash" bson:"enable_dash"`

	// 是否开通/启用 scm
	EnableScm bool `json:"enable_scm" bson:"enable_scm"`

	// 是否开通/启用 sale
	EnableSale bool `json:"enable_sale" bson:"enable_sale"`

	// 是否开通/启用 dish
	EnableDish bool `json:"enable_dish" bson:"enable_dish"`

	// 是否开通/启用 custome
	EnableCustome bool `json:"enable_custome" bson:"enable_custome"`

	// 是否开通/启用 sms
	EnableSms bool `json:"enable_sms" bson:"enable_sms"`

	// 是否开通/启用 join
	EnableJoin bool `json:"enable_join" bson:"enable_join"`

	// 是否开通/启用 system
	EnableSystem bool `json:"enable_system" bson:"enable_system"`

	// 是否开通/启用 store
	EnableStore bool `json:"enable_store" bson:"enable_store"`
}

// 根据用户角色分配查看的权限
// 1. 控制前段显示
// 2. 控制api的访问
type MenuDisplay struct {
	// 是否开通/启用 dash
	DisplayDash bool `json:"display_dash" bson:"display_dash"`
	// 是否开通/启用 scm
	DisplayScm bool `json:"display_scm" bson:"display_scm"`
	// 是否开通/启用 sale
	DisplaySale bool `json:"display_sale" bson:"display_sale"`
	// 是否开通/启用 dish
	DisplayDish bool `json:"display_dish" bson:"display_dish"`
	// 是否开通/启用 custome
	DisplayCustome bool `json:"display_custome" bson:"display_custome"`
	// 是否开通/启用 sms
	DisplaySms bool `json:"display_sms" bson:"display_sms"`
	// 是否开通/启用 join
	DisplayJoin bool `json:"display_join" bson:"display_join"`
	// 是否开通/启用 system
	DisplaySetting bool `json:"display_setting" bson:"display_setting"`
	// 是否开通/启用 stpre
	DisplayStore bool `json:"display_store" bson:"display_store"`
	// 是否开通/启用 report
	DisplayReport bool `json:"display_report" bson:"display_report"`
	// 是否开通/启用 online
	DisplayOnline bool `json:"display_online" bson:"display_online"`
	// 是否开通/启用 review
	DisplayReview bool `json:"display_review" bson:"display_review"`
	// 是否开通/启用 trade
	DisplayTrade bool `json:"display_trade" bson:"display_trade"`
}