package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

/**
 * @Author: cheney
 * @Date: 2022/11/14 11:11 PM
 * @Desc:
 */

type TomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Logo        string   `json:"logo,omitempty"`
	Navigation  []string `json:"navigation,omitempty"`
	Bilibili    string   `json:"bilibili,omitempty"`
	Avatar      string   `json:"avatar,omitempty"`
	UserName    string   `json:"userName,omitempty"`
	UserDesc    string   `json:"userDesc,omitempty"`
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *TomlConfig

func init() {
	Cfg = new(TomlConfig)
	var err error
	Cfg.System.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	Cfg.System.AppName = "ms-go-blog"
	Cfg.System.Version = 1.0
	_, err = toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err)
	}
}
