package gowebbase

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

func Config() Provider {
	if defaultConfig == nil {
		log.Panic("Please load config first...")
	}
	return defaultConfig
}

func LoadConfig() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	v := viper.New()


	v.SetDefault("port", 8000)
	v.SetDefault("VERSION", Version)
	v.SetConfigName(".env")
	v.SetConfigType("dotenv")
	v.AddConfigPath(".")
	v.ReadInConfig()
	v.AutomaticEnv()

	return v
}
