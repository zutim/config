package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件并返回配置对象
func LoadConfig(path []string) *viper.Viper {
	viperNew := viper.New()
	if len(path) == 0 {
		viperNew.SetConfigFile("config/db.yaml")
	} else {
		viperNew.SetConfigFile(path[0])
	}

	err := viperNew.ReadInConfig() // 查找并读取配置文件
	if err != nil {                // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if len(path) > 1 {
		for i := 1; i < len(path); i++ {
			viperNew.SetConfigFile(path[i])
			err = viperNew.MergeInConfig()
			if err != nil {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
		}
	}
	return viperNew
}
