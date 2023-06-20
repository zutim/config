package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := LoadConfig([]string{"db.yaml", "demo.json"})
	r := conf.Get("redis")
	if r != "127.0.0.1:30500" {
		t.Error("读取错误")
		return
	}

	d := conf.GetInt("otherWin")
	if d != 2000 {
		t.Error("读取demo.json错误")
		return
	}

	fmt.Println(r, d)
}
