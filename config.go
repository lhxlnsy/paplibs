package libs

import (
	"fmt"
	"path/filepath"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Debug  bool `env:"Debug" default:"true"`
	Logger struct {
		Enabled          bool     `env:"Enabled" default:"false"`
		Level            string   `env:"Level"`
		Test             string   `env:"Test"`
		Encoding         string   `env:"Encoding"`
		LogPath          string   `env:"LogPath"`
		OutputPaths      []string `env:"OutputPaths"`
		ErrorOutPutPaths []string `env:"ErrorOutputPaths"`
		EncoderConfig    struct {
			MessageKey   string `env:"MessageKey"`
			LevelKey     string `env:"LevelKey"`
			LevelEncoder string `env:"LevelEncoder"`
		}
	}
	DB struct {
		Prefix   string `env:"DBPrefix" default:"iris_"`
		Name     string `env:"DBName" default:""`
		Adapter  string `env:"DBAdapter" default:""`
		Host     string `env:"DBHost" default:""`
		Port     string `env:"DBPort" default:""`
		User     string `env:"DBUser" default:""`
		Password string `env:"DBPassword" default:""`
	}
	Redis struct {
		Host string `env:"RedisHost" default:"localhost"`
		Port string `env:"RedisPort" default:"6379"`
		Pwd  string `env:"RedisPwd" default:""`
	}
	Influxdb struct {
		DatabaseName string `env:"InfluxDBName" default:"planetarkpower"`
		Port         string `env:"InfluxDBPort" default:"8086"`
		Host         string `env:"InfluxDBHost" default:"localhost"`
		Precisions   string `env:"InfluxDBPrec" default:"ts"`
		Token        string `env:"InfluxDBToken" default:""`
		Bucket       string `env:"InfluxDBBucket" default:"TradingSystem"`
		Org          string `env:"InfluxDBOrg" default:"PlanetArkPower"`
	}
}{}

func InitConfig(p string) {

	configPath := filepath.Join(CWD(), "application.yml")
	if p != "" {
		configPath = p
	}
	_Logger.Print().Infow(fmt.Sprintf("Configuration File Path: %v", configPath))
	if err := configor.Load(&Config, configPath); err != nil {
		_Logger.Print().Error(fmt.Sprintf("Config Path:%s, Error:%s", configPath, err.Error()))
		return
	}
	if Config.Debug {
		_Logger.Print().Infow(fmt.Sprintf("配置项: %v", Config))
	}
}
