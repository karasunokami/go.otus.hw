package main

import (
	"flag"
	"fmt"
	"github.com/karasunokami/go.otus.hw/calendar/internal/calendar"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func init() {
	flag.String("configs", "configs/config.yaml", "path to yml configs file")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	configPath := viper.GetString("configs")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Couldn't read configuration file: %s", err.Error())
	}
}

func main() {
	var cld calendar.Calendar = calendar.NewCalendar()

	fmt.Println("Service started", cld)
}
