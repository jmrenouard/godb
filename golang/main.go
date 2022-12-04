package main

import (
	"fmt"
	"github.com/spf13/viper/remote"
	"github.com/spf13/viper"
	etcd "go.etcd.io/etcd/client"
)

func main() {
	viper.SetDefault("ETCD_PORT", "")
	viper.SetDefault("ETCD_ADDR", "Testing Endpoint")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println("ETCD Address is: ", viper.Get("ETCD_ADDR"))
	fmt.Println("ETCD Port is   : ", viper.Get("ETCD_PORT"))
	
}