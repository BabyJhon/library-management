package config

import "github.com/spf13/viper"

func InitConfig() error {
	viper.AddConfigPath("configs") //путь к файлу с конфигами
	viper.SetConfigName("config")  //файл config.yml
	return viper.ReadInConfig()    //считывает значения конфигов и записывает во внутренний объект viper'а
}
