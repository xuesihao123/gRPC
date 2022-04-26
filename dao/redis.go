package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RDb *redis.Client
)

func InitRedis()  error{
	RDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ ,err := RDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// redis operation hmset
//func SetMapForever(key string, field map[string]interface{}) (string, error) {
//	return RDb.HMSet(key, field).Result()
//}
//
//// redis operation hmget
//func GetMap(key string, fields ...string) ([]interface{}, error) {
//	return RDb.HMGet(key, fields...).Result()
//}
//
//// redis SADD
//func SetAdd(key string, field string) (int64, error) {
//	return RDb.SAdd(key, field).Result()
//}

// redis SADD
func SetKey(ctx context.Context,key string, field string,t time.Duration) (string, error) {
	return RDb.Set(ctx,key, field,t).Result()
}

// redis SADD
func GetKey(ctx context.Context,key string) (string, error) {
	return RDb.Get(ctx,key).Result()
}

//// redis SADD
//func ExpireKey(key string, t int) (error) {
//	return RDb.Expire(key, time.Duration(t)).Err()
//}
//
//// redis SISMEMBER
//func SetIsMember(key string, field string) (bool, error) {
//	return RDb.SIsMember(key, field).Result()
//}
//
//// redis SMEMBERS
//func GetSetMembers(key string) ([]string, error) {
//	return RDb.SMembers(key).Result()
//}
//func uniqueOccurrences(arr []int) bool {
//	 num := make(map[int]int)
//	 simple := make(map[int]int)
//	for _ , value := range arr{
//		num[value] = num[value]+1
//	}
//
//	for _ ,value := range num{
//		simple[value] = simple[value]+1
//		if simple[value] > 1{
//			return false
//		}
//	}
//	return true
//}
