package global

import (
	"fmt"
	"github.com/liuhongdi/digv14/pkg/setting"
	"time"
)
//服务器配置
type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
//数据库配置
type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
//日志配置
type LogSettingS struct {
	LogFilePath     string    //保存到的目录
	LogInfoFileName string    //info级日志文件的名字
	LogWarnFileName string    //warn级日志文件的名字
	LogAccessFileName string  //Access日志文件的名字
	LogFileExt      string    //文件的扩展名
}
//访问日志配置
type AccessLogSettingS struct {
	LogFilePath     string    //保存到的目录
	LogFileName string  //Access日志文件的名字
	LogFileExt      string    //文件的扩展名
}
//限流配置
type LimiterSettingS struct {
	CountPerSecond     int    //每秒的访问次数
}
//定义全局变量
var (
	ServerSetting   *ServerSettingS
	DatabaseSetting *DatabaseSettingS
	LogSetting *LogSettingS
	AccessLogSetting *AccessLogSettingS
	LimiterSetting *LimiterSettingS
)

//读取配置到全局变量
func SetupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Log", &LogSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("AccessLog", &AccessLogSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Limiter", &LimiterSetting)
	if err != nil {
		return err
	}

	fmt.Println("setting:")
	fmt.Println(ServerSetting)
	fmt.Println(DatabaseSetting)
	fmt.Println(LogSetting)
	fmt.Println(AccessLogSetting)
	fmt.Println(LimiterSetting)
	return nil
}

