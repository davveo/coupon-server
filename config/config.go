package config

import (
	"sync"

	"github.com/davveo/market-coupon/pkg/config"
	"github.com/davveo/market-coupon/pkg/config/log"
	"github.com/davveo/market-coupon/pkg/config/reader"
)

const (
	defaultServerHost = "0.0.0.0"
	defaultServerPort = 8080
)

// Config 配置项
type Config struct {
	AppName  string
	Host     string
	Port     int
	Meta     config.MetaEnv
	Redis    config.Redis
	MySql    config.Mysql
	Log      log.Log
	TestCron string
}

var (
	conf *Config
	once sync.Once
)

// Init 初始化配置
func Init() (*Config, error) {
	c := &Config{}
	r := reader.New()

	r.StringVar(&c.Host, "HOST", defaultServerHost)
	r.IntVar(&c.Port, "PORT", defaultServerPort)
	r.StringVar(&c.AppName, "APP_NAME", "")
	// meta
	r.StringVar(&c.Meta.Platform, "PLATFORM", "")
	r.StringVar(&c.Meta.Service, "SERVICE", "")
	r.StringVar(&c.Meta.Env, "ENV", "")
	r.StringVar(&c.Meta.Version, "VERSION", "")
	r.StringVar(&c.Meta.ID, "ID", "")
	// goods mysql
	r.StringVar(&c.MySql.DB, "MYSQL_DATABASE", "")
	r.StringVar(&c.MySql.Host, "MYSQL_HOST", "")
	r.IntVar(&c.MySql.Port, "MYSQL_PORT", 3306)
	r.StringVar(&c.MySql.User, "MYSQL_USER", "")
	r.StringVar(&c.MySql.Password, "MYSQL_PASSWORD", "")
	// redis
	r.StringVar(&c.Redis.Host, "REDIS_HOST", "localhost")
	r.IntVar(&c.Redis.Port, "REDIS_PORT", 6379)
	r.StringVar(&c.Redis.Password, "REDIS_PASSWORD", "")
	r.IntVar(&c.Redis.DB, "REDIS_DB", 0)
	// log
	r.StringVar(&c.Log.Path, "LOG_PATH", "")
	r.IntVar(&c.Log.MaxSize, "LOG_MAX_SIZE", 0)
	r.IntVar(&c.Log.MaxBackups, "LOG_MAX_BACKUPS", 0)
	r.IntVar(&c.Log.MaxAge, "LOG_MAX_AGE", 0)
	r.BoolVar(&c.Log.Compress, "LOG_COMPRESS", false)

	// cron
	r.StringVar(&c.TestCron, "CHECK_TABLE_DATA_EQUAL_CRON", "0 0 3 * * ?")

	once.Do(func() {
		conf = c
	})

	return c, nil
}

func Get() *Config {
	return conf
}
