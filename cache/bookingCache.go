package cache

import (
	"encoding/json"
	// "fmt"
	"time"

	"github.com/Aman130/HotelManagement/models"
	"github.com/go-redis/redis/v7"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key string, post []models.Hotel) {
	client := cache.getClient()

	// serialize Post object to JSON
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	client.Set(key, json, cache.expires*time.Second)
}

func (cache *RedisCache) Get(key string) []models.Hotel {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}

	post := []models.Hotel{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	// fmt.Println("In Get")
	// fmt.Println(post)
	return post
}



//For single data (not array)

// func (cache *RedisCache) Set(key string, post *models.Hotel) {
// 	client := cache.getClient()

// 	// serialize Post object to JSON
// 	json, err := json.Marshal(post)
// 	if err != nil {
// 		panic(err)
// 	}

// 	client.Set(key, json, cache.expires*time.Second)
// }

// func (cache *RedisCache) Get(key string) *models.Hotel {
// 	client := cache.getClient()

// 	val, err := client.Get(key).Result()
// 	if err != nil {
// 		return nil
// 	}

// 	post := models.Hotel{}
// 	err = json.Unmarshal([]byte(val), &post)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &post
// }

