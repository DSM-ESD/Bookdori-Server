package database

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// Database 구조체는 DB에 대한 정보를 담습니다.
type Database struct {
	Name string
	DB   *mongo.Database
}

// C 메서드는 해당하는 DB의 Collection을 반환합니다
func (db *Database) C(c string) *mongo.Collection {
	if db.DB == nil {
		log.Println("DB is nil")
		return nil
	}
	return db.DB.Collection(c)
}