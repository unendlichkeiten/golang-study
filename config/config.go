package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerSettingS struct {
	RunMode         string        `json:"RunMode,omitempty" yaml:"RunMode,omitempty"`
	HttpPort        string        `json:"HttpPort,omitempty" yaml:"HttpPort,omitempty"`
	ReadTimeout     time.Duration `json:"ReadTimeout,omitempty" yaml:"ReadTimeout,omitempty"`
	WriteTimeout    time.Duration `json:"WriteTimeout,omitempty" yaml:"WriteTimeout,omitempty"`
}

type AppSettingS struct {
	DefaultPageSize int32         `json:"DefaultPageSize,omitempty" yaml:"DefaultPageSize,omitempty"`
	MaxPageSize     int32         `json:"MaxPageSize,omitempty" yaml:"MaxPageSize,omitempty"`
	LogSavePath     string        `json:"LogSavePath,omitempty" yaml:"LogSavePath,omitempty"`
	LogFileName     string        `json:"LogFileName,omitempty" yaml:"LogFileName,omitempty"`
	LogFileExt      string        `json:"LogFileExt,omitempty" yaml:"LogFileExt,omitempty"`
}

type DatabaseSettingS struct {
	DBType          string        `json:"DBType,omitempty" yaml:"DBType,omitempty"`
	Username        string        `json:"Username,omitempty" yaml:"Username,omitempty"`
	Password        string        `json:"Password,omitempty" yaml:"Password,omitempty"`
	Host            string        `json:"Host,omitempty" yaml:"Host,omitempty"`
	DBName          string        `json:"DBName,omitempty" yaml:"DBName,omitempty"`
	TablePrefix     string        `json:"TablePrefix,omitempty" yaml:"TablePrefix,omitempty"`
	Charset         string        `json:"Charset,omitempty" yaml:"Charset,omitempty"`
	ParseTime       bool          `json:"ParseTime,omitempty" yaml:"ParseTime,omitempty"`
	MaxIdleConns    int32         `json:"MaxIdleConns,omitempty" yaml:"MaxIdleConns,omitempty"`
	MaxOpenConns    int32         `json:"MaxOpenConns,omitempty" yaml:"MaxOpenConns,omitempty"`

}

type Config struct{
	vp *viper.Viper
}

var _configs = make(map[string]interface{})

var (
	ServerSetting   *ServerSettingS
	AppSetting      *AppSettingS
	DatabaseSetting *DatabaseSettingS
)

func NewConfig(configs ...string) (*Config, error) {
	vp := viper.New()
	for _, config := range configs {
		if config != "" {
			vp.SetConfigName(config)
			vp.AddConfigPath(".")
		}
	}

	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	s := &Config{vp: vp}

	s.WatchConfig()

	return s, nil
}

func (c *Config) WatchConfig() {
	go func() {
		c.vp.WatchConfig()
		c.vp.OnConfigChange(func(in fsnotify.Event) {
			c.ReloadConfigs()
		})
	}()
}

func (c *Config) ReloadConfigs() error {
	for k, v := range _configs {
		if err := c.ReadConfig(k, v); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) ReadConfig(k string, v interface{}) error {
	if err := c.vp.UnmarshalKey(k, v); err != nil {
		return err
	}

	_configs[k] = v

	return nil
}
