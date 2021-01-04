package global

import (
	"github.com/liuhongdi/digv14/pkg/zaplog"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

//创建logger
func SetupLogger() (error) {
	var err error
	filepath:= LogSetting.LogFilePath
	infofilename:= LogSetting.LogInfoFileName
	warnfilename:= LogSetting.LogWarnFileName
	fileext:= LogSetting.LogFileExt

	//infofilename,warnfilename,fileext string
    Logger,err = zaplog.GetInitLogger(filepath,infofilename,warnfilename,fileext)

	if err != nil {
		return err
	}
	defer Logger.Sync()
    return nil
}


