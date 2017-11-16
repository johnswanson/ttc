package main

import (
	"fmt"
	"github.com/johnswanson/ttc"
	"github.com/johnswanson/ttc/api"
	"github.com/johnswanson/ttc/dialog"
	"github.com/johnswanson/ttc/pings"
	"github.com/spf13/viper"
	"os"
	"time"
)

func configure() {
	viper.SetConfigName("ttc.conf")
	viper.AddConfigPath("$HOME/.config/ttc")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func main() {
	configure()

	API := ttc.API{
		URL:   viper.Get("URL").(string),
		Token: viper.Get("token").(string),
	}

	config := ttc.Config{}
	err := api.GetConfig(API, &config)
	if err != nil {
		fmt.Printf("Failed to retrieve configuration!: %v", err)
		os.Exit(1)
	}
	pings := pings.PingChannel(config)
	fmt.Printf("CONFIG: %v\n", config)

	ticker := time.NewTicker(time.Second)
	nextPing := <-pings
	for t := range ticker.C {
		now := t.Unix()
		for nextPing < now {
			if now-nextPing <= 60*3 {
				err, ping := dialog.Request(nextPing)
				p := &ping
				if err == nil {
					fmt.Printf("Saving!\n")
					api.Save(API, p)
				} else {
					fmt.Printf("ERROR: %v\n", err)
				}
			}
			nextPing = <-pings
		}
	}
}
