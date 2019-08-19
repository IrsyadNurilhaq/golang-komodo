package main

import (
	"fmt"
	route "komodo/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func readConfig() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		viper.SetEnvPrefix("hr-api")
		viper.AllowEmptyEnv(true)
		viper.AutomaticEnv()
	}
}

func main() {
	readConfig()
	r := route.Routes()
	r.Run(":8888")
}
