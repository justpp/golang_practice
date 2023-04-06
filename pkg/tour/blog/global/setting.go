package global

import (
	"giao/pkg/tour/blog/pkg/logger"
	"giao/pkg/tour/blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
	EmailSetting    *setting.EmailSettingS
)
