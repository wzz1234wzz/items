package config

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"runtime"
	"strings"
)

// CustomT CustomT
type Mysql struct {
	DNS string `yaml:"dns"`
}

type SqlServer struct {
	DNS string `yaml:"dns"`
}

// CustomT CustomT
type CustomT struct {
	Mysql     Mysql     `yaml:"mySql"`
	SqlServer SqlServer `yaml:"sqlServer"`
}

// Custom Custom
var Custom CustomT

// ReadConfig ReadConfig for custom
func ReadConfig(configName string, configPath string, configType string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(configName)
	v.AddConfigPath(configPath)
	v.SetConfigType(configType)
	err := v.ReadInConfig()
	if err != nil {
		return nil
	}
	return v
}

func CurrentFileDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "失败"
	}
	i := strings.LastIndex(file, "/")
	if i < 0 {
		i = strings.LastIndex(file, "\\")
	}

	return string(file[0 : i+1])
}

// InitConfig InitConfig
func InitConfig() {
	path := CurrentFileDir()
	v := ReadConfig("dev_custom", path, "yml")
	md := mapstructure.Metadata{}
	err := v.Unmarshal(&Custom, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.Metadata = &md
	})
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("InitConfig Success!")
}
