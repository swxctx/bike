package cfg

import (
	"io/ioutil"
	"os"

	"github.com/domego/ginkits/db"
	"github.com/domego/ginkits/redis"
	"github.com/domego/gokits/log"

	orm "github.com/swxctx/bike/model/bike"

	yaml "gopkg.in/yaml.v2"
)

var (
	// ConfigFile config 文件地址
	ConfigFile = "config/config.yaml"
	// Conf config 实例
	Conf *Config
)

// Config config struct
type Config struct {
	Address  string                 `yaml:"address"`
	Env      string                 `yaml:"env"`
	LogLevel string                 `yaml:"log_level"`
	Redis    *rediskits.RedisConfig `yaml:"redis"`
	DB       struct {
		Bike *db.DBConfig `yaml:"bike"`
	} `yaml:"db"`
}

// readConfig read and parse configFile
func ReadConfig(exit bool) *Config {
	bs, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Errorf("load config file failed, the path is %s", ConfigFile)
		if exit {
			os.Exit(1)
		}
	}
	config := &Config{}
	err = yaml.Unmarshal(bs, config)
	if err != nil {
		log.Errorf("unmarshal config failed, %s", err)
		if exit {
			os.Exit(1)
		}
	}
	return config
}

func Init() *Config {
	Conf = ReadConfig(true)
	InitConfig()
	return Conf
}

func Reload() {
	newConfig := ReadConfig(false)
	if newConfig != nil {
		log.Infof("reload config")
		Conf = newConfig
		InitConfig()
	}
}

func InitConfig() {
	log.DefaultLogger.Level = log.NewLevel(Conf.LogLevel)
	if len(Conf.Env) == 0 {
		Conf.Env = "debug"
	}
	orm.DBConfig = Conf.DB.Bike
	if Conf.Redis != nil {
		redisClient := rediskits.NewClient(Conf.Redis)
		orm.RedisClient = redisClient
	}
}
