package config

// RocketMQ rocket mq 配置
type RocketMQ struct {
	Endpoint         string
	AccessKey        string
	SecretKey        string
	InstanceId       string
	DefaultGroupId   string
	DefaultGroupIdV2 string
	DefaultGroupIdV3 string
}
