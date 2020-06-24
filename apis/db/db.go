package db

import (
	"bookdori-server/database"
)

// MongoDB 변수는 MongoDB에 대한 정보를 담은 구조체 타입의 변수입니다
var MongoDB, _ = database.NewMongoDB("mongodb://127.0.0.1:27017")