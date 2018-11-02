package config

import (
	"os"
	"strconv"
)

// Config is an app's configuration
type Config struct {
	Port                   int
	DBURL                  string
	EncryptionKey          string
	Secret                 string
	AccessTokenExpiryTime  int
	RefreshTokenExpiryTime int
	AuthURL                string
	ConsumerSecret         string
}

// New returns config
func New() *Config {
	cfg := &Config{}
	cfg.Port = getEnvInt("PORT")
	cfg.DBURL = getEnvStr("DB_URL")
	cfg.EncryptionKey = getEnvStr("ENCRYPTION_KEY")
	cfg.Secret = getEnvStr("SECRET")
	cfg.AccessTokenExpiryTime = getEnvInt("ACCESS_TOKEN_EXPIRY_TIME")
	cfg.RefreshTokenExpiryTime = getEnvInt("REFRESH_TOKEN_EXPIRY_TIME")
	cfg.AuthURL = getEnvStr("AUTH_URL")
	cfg.ConsumerSecret = getEnvStr("CONSUMER_SECRET")
	return cfg
}

func getEnvStr(key string) string {
	return os.Getenv(key)
}

func getEnvInt(key string) int {
	val, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return val
}
