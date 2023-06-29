package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("dbconfig")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration file: %s", err))
	}

	err = viper.Unmarshal(&dbconfig)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration: %s", err))
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	error := DBConnection.Connect()
	if error != nil {
		fmt.Println(error)
	}
	StartGrpcServer()
}
