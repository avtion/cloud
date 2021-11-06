package main

import "github.com/spf13/viper"

const (
	defaultConfigFilename = "config"
	defaultConfigType     = "yaml"
	defaultConfigPath     = "."
)

// loadWorkDirConfig 从工作目录加载配置
func loadWorkDirConfig() (*globalConfig, error) {
	viper.SetConfigName(defaultConfigFilename)
	viper.SetConfigType(defaultConfigType)
	viper.AddConfigPath(defaultConfigPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	fileConfig := new(globalConfig)
	if err := viper.Unmarshal(&fileConfig); err != nil {
		return nil, err
	}
	return fileConfig, nil
}
