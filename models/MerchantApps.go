package models

import (
	"context"

	"github.com/r2day/db"
	"go.mongodb.org/mongo-driver/bson"
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

	// 应用唯一标识
	Key string `json:"key" bson:"key"`

	// 应用介绍
	Desc string `json:"desc"`

	// 应用配置
	// 用于指导前端显示sidebar
	AppMenu ApplicationMenu `json:"application_menu" bson:"application_menu"`
}

// MerchantApps 通过id获取对象
func (m * MerchantApps) GetOneById(id string) (*MerchantApps, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	objID, _ := primitive.ObjectIDFromHex(id)
	coll := db.MDB.Collection(MerchantAppConfCollection)
	err := coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objID}}).Decode(m)

	if err != nil {
		return nil,  err
	}
	return m, nil
}