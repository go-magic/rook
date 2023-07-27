package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

type Server struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func TestName(t *testing.T) {
	mongo := NewMongo("mongodb://192.168.209.129:27017")
	mongo.SetHeartTime(time.Second * 10)
	if err := mongo.Init(); err != nil {
		t.Fatal(err)
	}
	db := mongo.client.Database("aliyun")
	collection := db.Collection("servers")
	//filter := Server{Name: "飞天"}
	filter := bson.M{"version": "1.0.1"}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		t.Fatal(err)
	}
	results := make([]*Server, 0)
	if err = cur.All(context.Background(), &results); err != nil {
		t.Fatal(err)
	}
	t.Log("测试通过")
}
