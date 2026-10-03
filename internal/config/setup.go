package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client
var Ctx = context.Background()

func InitDependencies() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek DB: ", err)
	}
	DB = db
	fmt.Println("✅ PostgreSQL Terhubung!")

	RedisClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})
	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		log.Fatal("Gagal konek Redis: ", err)
	}
	fmt.Println("✅ Redis Terhubung!")
}
