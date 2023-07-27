package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

type Mongo struct {
	dbUrl     string
	dbName    string
	heartTime time.Duration
	client    *mongo.Client
	lock      sync.Mutex
}

func NewMongo(dbUrl string) *Mongo {
	return &Mongo{
		dbUrl: dbUrl,
	}
}

func (m *Mongo) SetHeartTime(heartTime time.Duration) {
	m.heartTime = heartTime
}

func (m *Mongo) Init() error {
	err := m.dail()
	go m.ping()
	return err
}

func (m *Mongo) dail() error {
	m.lock.Lock()
	var err error
	clientOptions := options.Client().ApplyURI(m.dbUrl)
	// Connect to MongoDB 连接数据库,返回客户端对象client
	m.client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	m.lock.Unlock()
	return err
}

func (m *Mongo) heart() {
	ticker := time.NewTicker(m.heartTime)
	for {
		select {
		case <-ticker.C:
			m.ping()
			ticker.Reset(m.heartTime)
		}
	}
}

func (m *Mongo) ping() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	if err := m.client.Ping(ctx, nil); err != nil {
		if err = m.dail(); err != nil {

		}
	}
}

func (m *Mongo) Close() error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	return m.client.Disconnect(ctx)
}
