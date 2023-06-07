package config

// Redis redis配置
type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}
