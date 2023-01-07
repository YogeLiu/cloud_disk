package conf

import (
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml/v2"

	"github.com/YogeLiu/CloudDisk/pkg/util"
)

type redis struct {
	Network  string
	Addr     string
	Password string
	DB       int
}

type database struct {
	Type       string
	User       string
	Password   string
	Host       string
	Name       string
	DBFile     string
	Port       int
	Charset    string
	UnixSocket bool
}

// system 系统通用配置
type system struct {
	Secret     string
	HashIDSalt string
	Listen     string
	Debug      bool
}

func Init(path string) {
	if path == "" || !util.PathExsit(path) {
		confContent := util.Replace(map[string]string{
			"{Secret}":     util.RandStringRunes(64),
			"{HashIDSalt}": util.RandStringRunes(64),
		}, defaultConf)
		path = "./conf.toml"
		file, err := os.Create(path)
		if err != nil {
			util.Log().Panic("Failed to create config file: %s", err)
		}
		defer file.Close()
		_, err = file.WriteString(confContent)
		if err != nil {
			util.Log().Panic("Failed to write config file: %s", err)
		}
	}
	config := map[string]interface{}{
		"Database": DatabaseConfig,
		"System":   SystemConfig,
		"Redis":    RedisConfig,
	}
	file, err := os.Open(path)
	if err != nil {
		util.Log().Panic(err.Error())
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		util.Log().Panic(err.Error())
	}
	err = toml.Unmarshal(content, &config)
	if err != nil {
		util.Log().Panic(err.Error())
	}
}
