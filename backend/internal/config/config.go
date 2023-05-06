package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
)

// 配置实例，单例
var (
	Config         *_config
	ConfigFileName string = filepath.Join(file.SelfDir(), "ticket.yaml")
)

type _system struct {
	DbFile  string `yaml:"db_file"`
	BaseDir string `yaml:"base_dir"`
	Mode    string `yaml:"mode"`
	IsDemo  bool   `yaml:"is_demo"`
	Port    int    `yaml:"port"`
}
type _log struct {
	Level     string `yaml:"level"`
	TimeZone  string `yaml:"time_zone"`
	LogName   string `yaml:"log_name"`
	LogSuffix string `yaml:"log_suffix"`
	MaxBackup int    `yaml:"max_backup"`
}

type _config struct {
	System _system `yaml:"system"`
	Log    _log    `yaml:"log"`
}

// 读取配置文件
func Read() error {

	// fmt.Printf("configFileName: %v\n", ConfigFileName)

	b, err := os.ReadFile(ConfigFileName)
	if err != nil {
		return err
	}

	err2 := yaml.Unmarshal(b, &Config)
	if err2 != nil {
		return err2
	}

	return nil
}

// 新建配置文件
func New() error {
	conf := &_config{
		System: _system{
			DbFile:  "ticket.dbf",
			BaseDir: file.SelfDir(),
			Mode:    "dev",
			IsDemo:  true,
			Port:    9999,
		},
		Log: _log{
			Level:     "debug",
			TimeZone:  "Asia/Shanghai",
			LogName:   "ticket",
			LogSuffix: ".log",
			MaxBackup: 10,
		},
	}
	Config = conf
	return nil
}

// 保存配置文件
func Save() error {

	data, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}

	err2 := os.WriteFile(ConfigFileName, data, 0655)
	if err2 != nil {
		return err2
	}
	return nil
}

// 显示配置
func Show() error {
	out, err := yaml.Marshal(Config)
	if err != nil {
		return err
	}

	fmt.Printf(`configurations:
%s`, string(out))
	return nil
}
