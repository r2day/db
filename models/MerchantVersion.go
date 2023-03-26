package models

import (
	"context"
	"fmt"

	"github.com/r2day/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	// 包含应用列表
	ApplicationList []string `json:"application_list" bson:"application_list"`

}

// type ApplicationShortcut struct {
// 	// 标签
// 	Badge string `json:"badge"`

// 	// 应用价格
// 	Price uint `json:"price" bson:"price"`

// 	// 版本介绍
// 	Desc string `json:"desc"`
// }


// MerchantVersion 通过id获取对象
func (m * MerchantVersion) GetOneById(id string) (*MerchantVersion, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	objID, _ := primitive.ObjectIDFromHex(id)
	coll := db.MDB.Collection(MerchantVersionCollection)
	err := coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objID}}).Decode(m)

	if err != nil {
		return nil,  err
	}
	return m, nil
}

// MerchantVersion 通过id获取对象
func (m * MerchantVersion) GetApplicationsById(id string) ([]*ApplicationConfig, error) {
	// TODO result using custom struct instead of bson.M
	// because you should avoid to export something to customers
	objID, _ := primitive.ObjectIDFromHex(id)
	coll := db.MDB.Collection(MerchantVersionCollection)
	err := coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objID}}).Decode(m)

	if err != nil {
		return nil,  err
	}
	applicationList := make([]*ApplicationConfig, 0)
	roleManageColl := db.MDB.Collection(ApplicationConfigCollectionName)

	applicationMongoIdList := make([]primitive.ObjectID, 0)
	for _, i := range m.ApplicationList {
		objID, _ := primitive.ObjectIDFromHex(i)
		applicationMongoIdList = append(applicationMongoIdList, objID)

	}

	filter := bson.M{"_id": bson.M{"$in": applicationMongoIdList}}
	cursor, err := roleManageColl.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &applicationList)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("no application was found")
		return nil, err
	}
	if err != nil {
		fmt.Printf("query application by ids in failed")
		return nil, err
	}
	return applicationList, nil
}