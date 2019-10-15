package lib


type BaseConf struct {
	DebugMode string `mapstructure:"debug_mode"`
	TimeLocation string `mapstructure:"time_location"`
	Log	LogConfig `mapstructure:"log"`
	Base struct {
		DebugMode string `mapstructure:"debug_mode"`
		TimeLocation string `mapstructure:"time_location"`
	}`mapstructure:"base"`
}

type LogConfFileWriter struct {
	On bool `mapstructure:"on"`
	LogPath string `mapstructure:"log_path"`
	RotateLogPath string `mapstructure:"rotate_log_path"`
	WfLogPath string `mapstructure:"wf_log_path"`
	RotateWfLogPath string `mapstructure:"rotate_wf_path"`
}

type LogConfConsoleWriter struct {
	On bool `mapstructure:"on"`
	Color bool `mapstructure:"color"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`

}