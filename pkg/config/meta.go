package config

type AppEnv int8

const (
	EnvDev  = "dev"
	EnvTest = "test"
	EnvPre  = "pre"
	EnvProd = "prod"
)

// 定义应用基本信息
type MetaEnv struct {
	// 应用所属业务平台
	Platform string
	// 应用所属服务
	Service string
	// 运行环境: dev/test/pre/prod
	Env string
	// 版本号
	Version string
	//id
	ID string
}
