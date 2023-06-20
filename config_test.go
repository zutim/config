package config

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	LoadConfig([]string{"db.yaml", "demo.json"})
	r := viper.Get("redis")
	if r != "127.0.0.1:30500" {
		t.Error("读取错误")
		return
	}

	d := viper.GetInt("otherWin")
	if d != 2000 {
		t.Error("读取demo.json错误")
		return
	}

	fmt.Println(r, d)
}
