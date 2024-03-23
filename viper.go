package core

import (
	"bytes"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperOption func(*viper.Viper)

func NewViper(config any, options ...ViperOption) *viper.Viper {
	v := viper.New()
	for _, option := range options {
		option(v)
	}
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("fatal error unmarshal config: %s \n", err)
	}

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			log.Fatalf("fatal error unmarshal config: %s \n", err)
		}
	})
	v.WatchConfig()

	v.AutomaticEnv()
	return v
}

func ViperSetEnvPrefix(in string) ViperOption {
	return func(v *viper.Viper) {
		v.SetEnvPrefix(in)
	}
}

func ViperSetConfigFile(in string) ViperOption {
	return func(v *viper.Viper) {
		v.SetConfigFile(in)
		if err := v.ReadInConfig(); err != nil {
			log.Fatalf("fatal error config file: %s \n", err)
		}
	}
}

func ViperReadConfig(t string, c []byte) ViperOption {
	return func(v *viper.Viper) {
		v.SetConfigType(t)
		v.ReadConfig(bytes.NewBuffer(c))
	}
}
