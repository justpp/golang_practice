package global

import (
	"giao/tour/blog/pkg/logger"
	"giao/tour/blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
