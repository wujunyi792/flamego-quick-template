package types

import "time"

type Cache interface {
	GetInt(key string) (int, bool)
	GetInt64(key string) (int64, bool)
	GetFloat32(key string) (float32, bool)
	GetFloat64(key string) (float64, bool)
	GetString(key string) (string, bool)
	GetBool(key string) (bool, bool)
	Set(Key string, value any, expireDuration time.Duration) error
	Del(key string) bool
}
