package models

import (
	"context"

	"github.com/r2day/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ApplicationConfigCollectionName = "application_config_collection"
)

//     {
// 	"sub_menu": [
// 		{
// 			"to": "/scm/itemCategory",
// 			"state": {
// 				"_scrollToTop": true,
// 			},
// 			"primary_text": "pos.submenu.scm.itemCategory",
// 			"left_icon": "test",
// 			"dense": false,
// 		}
// 	],
// },
// ApplicationConfig 应用配置 
type ApplicationConfig struct {
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
	// 例如: scm, product, order,
	Name string `json:"name"`
	// 功能标识
	Logo string `json:"logo"`
	// 应用列表
	AppMenuList []*ApplicationMenu  `json:"app_menu_list" bson:"app_menu_list"`
	
	// 应用state
	AppStateMap map[string]bool `json:"app_state_map" bson:"app_state_map"`
}

// ApplicationMenu 应用菜单
// 	"is_default_open": false,
// 	"name": "pos.menu.scm",
// 	"icon": "settings",
// 	"dense": false,
// 	"enable": true,
// 	"display": true,
// 	"toggle": "settings",
type ApplicationMenu struct {
	// 是否默认展开
	IsDefaultOpen bool `json:"is_default_open" bson:"is_default_open"`
	// 功能名称，格式为 xx.yy.zz 
	// 便于i18国际化处理
	Name string `json:"name"`
	// 图标名称
	// 前端通过名字找到对应的图标
	Icon string `json:"icon"`
	// Dense
	Dense bool `json:"dense" bson:"dense"`
	// enable 是否启用
	Enable bool `json:"enable" bson:"enable"`
	// display 是否显示
	Display bool `json:"display" bson:"display"`
	// 是否选中
	Toggle string `json:"toggle" bson:"toggle"`
	// 子目录
	AppSubMenuList []*ApplicationSubMenu `json:"app_sub_menu_list" bson:"app_sub_menu_list"`
}

// 

// 		{
// 			"to": "/scm/itemCategory",
// 			"state": {
// 				"_scrollToTop": true,
// 			},
// 			"primary_text": "pos.submenu.scm.itemCategory",
// 			"left_icon": "test",
// 			"dense": false,
// 		}
type ApplicationSubMenu struct {
	// 跳转页面
	To string `json:"to"`
	// 初级文字
	PrimaryText string `json:"primary_text" bson:"primary_text"`
	// LeftIcon 图标名称
	LeftIcon string `json:"left_icon" bson:"left_icon"`
	// Dense
	Dense bool `json:"dense" bson:"dense"`
}


// GetByMerchantId 通过商户id获取对象
func (m * ApplicationConfig) GetByMerchantId(merchantId string) (*ApplicationConfig, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	coll := db.MDB.Collection(ApplicationConfigCollectionName)
	err := coll.FindOne(context.TODO(),
		bson.D{
			{Key: "merchant_id", Value: merchantId},
			{Key: "status", Value: true},
		}).Decode(m)
	if err != nil {
		return nil,  err
	}
	return m, nil
}