package cfg

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Bot BotCFG `yaml:"bot"`
}

type BotCFG struct {
	Token    string `yaml:"token"`
	Admin    int64  `yaml:"admin"`
	Group    int64  `yaml:"group"`
	TimeZone int8   `yaml:"time_zone"`
}

var Get *Config
var onceCfg sync.Once

// Load загружает конфигурацию базы данных, и конфигурации для бота.
func Load() {
	onceCfg.Do(func() {
		Get = &Config{}

		if err := cleanenv.ReadConfig("configs/config.yml", Get); err != nil {
			log.Print(err)
		}
	})
}
