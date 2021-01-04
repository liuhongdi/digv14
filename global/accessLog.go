package global

import (
	"github.com/liuhongdi/digv14/pkg/zaplog"
	"go.uber.org/zap"
)

var (
	AccessLogger *zap.SugaredLogger
)

//创建logger
func SetupAccessLogger() (error) {
	var err error
	filepath:= AccessLogSetting.LogFilePath
	filename:= AccessLogSetting.LogFileName
	//warnfilename:= LogSetting.LogWarnFileName
	fileext:= AccessLogSetting.LogFileExt

	//infofilename,warnfilename,fileext string
	AccessLogger,err = zaplog.GetInitAccessLogger(filepath,filename,fileext)

	if err != nil {
		return err
	}
	defer AccessLogger.Sync()
    return nil
}


