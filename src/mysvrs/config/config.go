package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/nybuxtsui/log"
)

type ServerConfig struct {
	Http  string
	Https string
}

var SConfig ServerConfig

func Init(configFile string) {
	tomlstr, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error("read file error: %s", err)
	}
	if _, err := toml.Decode(string(tomlstr), &SConfig); err != nil {
		log.Error("toml.Decode error: %s", err)
	}
	log.Debug("%#v", SConfig)
}
