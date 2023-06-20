package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置文件并返回配置对象
func LoadConfig(path []string) {
	if len(path) == 0 {
		viper.SetConfigFile("config/db.yaml")
	} else {
		viper.SetConfigFile(path[0])
	}

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if len(path) > 1 {
		for i := 1; i < len(path); i++ {
			viper.SetConfigFile(path[i])
			err = viper.MergeInConfig()
			if err != nil {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
		}
	}
}
