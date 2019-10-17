package lib

import (
	"database/sql"
	"github.com/spf13/viper"
)

type BaseConf struct {
	DebugMode    string    `mapstructure:"debug_mode"`
	TimeLocation string    `mapstructure:"time_location"`
	Log          LogConfig `mapstructure:"log"`
	Base         struct {
		DebugMode    string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	} `mapstructure:"base"`
}

type LogConfFileWriter struct {
	On              bool   `mapstructure:"on"`
	LogPath         string `mapstructure:"log_path"`
	RotateLogPath   string `mapstructure:"rotate_log_path"`
	WfLogPath       string `mapstructure:"wf_log_path"`
	RotateWfLogPath string `mapstructure:"rotate_wf_path"`
}

type LogConfConsoleWriter struct {
	On    bool `mapstructure:"on"`
	Color bool `mapstructure:"color"`
}

type LogConfig struct {
	Level string               `mapstructure:"log_level"`
	FW    LogConfFileWriter    `mapstructure:"file_writer"`
	CW    LogConfConsoleWriter `mapstructure:"console_writer"`
}

// mysql
type MysqlMapConf struct {
	List map[string]*MysqlMapConf `mapstructure:"list"`
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleCoon     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

// redis
type RedisMapConf struct {
	List map[string]*RedisMapConf `mapstructure:"list"`
}

type RedisConf struct {
	ProxyList    []string `mapstructure:"proxy_list"`
	ConnTimeout  int      `mapstructure:"conn_timeout"`
	ReadTimeout  int      `mapstructure:"read_timeout"`
	WriteTimeout int      `mapstructure:"write_timeout"`
}

var (
	ConfBase        *BaseConf
	DBMapPool       map[string]*sql.DB
	GORMMapPool     map[string]*gorm.DB
	DBDefaultPool   *sql.DB
	GORMDefaultPool *gorm.DB
	ConfRedis       *RedisConf
	ConfRedisMap    *RedisMapConf
	ViperConfMap    map[string]*viper.Viper
)

func GetBaseConf() *BaseConf {
	return ConfBase
}

func InitBaseConf(path string) error {
	ConfBase = &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}

	if ConfBase.DebugMode == "" {
		if ConfBase.Base.DebugMode != "" {
			ConfBase.DebugMode = ConfBase.Base.DebugMode
		} else {
			ConfBase.DebugMode = "debug"
		}
	}

	if ConfBase.TimeLocation == "" {
		if ConfBase.Base.TimeLocation != "" {
			ConfBase.TimeLocation = ConfBase.Base.TimeLocation
		} else {
			ConfBase.TimeLocation = "Asia/Shanghai"
		}
	}
	if ConfBase.Log.Level == "" {
		ConfBase.Log.Level = "trace"
	}

	logConf := dLog.LogConfig{
		Level: ConfBase.Log.Level,

		FW: dLog.ConfFileWriter{
			On:              ConfBase.Log.FW.On,
			LogPath:         ConfBase.Log.FW.LogPath,
			RotateLogPath:   ConfBase.Log.FW.RotateLogPath,
			WfLogPath:       ConfBase.Log.FW.WfLogPath,
			RotateWfLogPath: ConfBase.Log.FW.RotateWfLogPath,
		},
		CW: dLog.ConfConsoleWriter{
			On:    ConfBase.Log.CW.On,
			Color: ConfBase.Log.CW.Color,
		},
	}
	if err := dLog.SetupDefaultLogWithConf(logConf); err != nil {
		panic(err)
	}
	dLog.SetLayout("2006-01-02T15:04:05.000")
	return nil
}

func InitRedisConf(path string) error {
	ConfRedis := &RedisMapConf{}
	err := ParseConfig(path, ConfRedis)
	if err != nil {
		return err
	}
	ConfRedisMap = ConfRedis
	return nil
}
