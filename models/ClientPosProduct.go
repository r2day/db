package models


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientPosProductModel struct {
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
	// 分类名称
	Name string `json:"name"`
	// 单价
	Price float64 `json:"price"  bson:"price"`
	// 库存
	Sku float64 `json:"sku"  bson:"sku"`
	// 图片
	Icon string `json:"icon"  bson:"icon"`
	// 跳转链接
	Link string `json:"link"  bson:"link"`
}
