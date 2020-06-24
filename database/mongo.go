package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB 구조체는 몽고DB에 대한 정보를 담습니다
type MongoDB struct {
	Client    *mongo.Client
	Databases map[string]*Database
	ctx       context.Context
	cancel    context.CancelFunc
}

// Connect 메서드는 몽고DB에 연결하니다
func (m *MongoDB) Connect() error {
	if err := m.Client.Connect(m.ctx); err != nil {
		return err
	}
	return nil
}

// Disconnect 메서드는 몽고DB 연결을 끊습니다.
func (m *MongoDB) Disconnect() {
	m.Client.Disconnect(m.ctx)
}

// Ping 메서드는 몽고DB 연결을 확인합니다.
func (m *MongoDB) Ping() error {
	if err := m.Client.Ping(m.ctx, readpref.Primary()); err != nil {
		return err
	}
	return nil
}

// NewClient 메서드는 몽고DB Client를 생성합니다.
func (m *MongoDB) NewClient(addr string) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		return err
	}
	m.Client = client
	return nil
}

// DB 메서드는 해당하는 DB를 반환합니다
func (m *MongoDB) DB(db string) *Database {
	if m.Client == nil {
		log.Println("Client is nil")
		return nil
	}
	if _, ok := m.Databases[db]; !ok {
		m.Databases[db] = new(Database)
		m.Databases[db].Name = db
		m.Databases[db].DB = m.Client.Database(db)
	}
	return m.Databases[db]
}

// NewMongoDB 메서드는 MongoDB 객체를 생성합니다.
func NewMongoDB(addr string) (*MongoDB, error) {
	m := &MongoDB{}
	if err := m.NewClient(addr); err != nil {
		return nil, err
	}

	m.Databases = make(map[string]*Database)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m.ctx = ctx
	m.cancel = cancel

	if err := m.Connect(); err != nil {
		return nil, err
	}

	if err := m.Ping(); err != nil {
		return nil, err
	}
	return m, nil
}