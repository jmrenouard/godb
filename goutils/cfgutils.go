package main

import (
	"github.com/spf13/viper"
)

func ReadConfigWithDefaults(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	var defaultValues map[string]interface{}

	if defaults == nil {
		defaultValues = map[string]interface{}{}
	} else {
		defaultValues = defaults
	}

	v := viper.New()
	for key, value := range defaultValues {
		v.SetDefault(key, value)
	}
	if filename != "none" {
		v.SetConfigName(filename)
	}
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}

func ReadConfig(filename string) *viper.Viper {
	v, err := ReadConfigWithDefaults(filename, nil)
	if err != nil {
		Panic(err.Error())
	}
	return v
}

func NewConfig() *viper.Viper {
	v := viper.New()
	return v
}
