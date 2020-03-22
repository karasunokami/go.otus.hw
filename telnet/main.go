package main

import (
	"errors"
	"fmt"
	"github.com/karasunokami/go.otus.hw/telnet/internal/client"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

var (
	timeout int
	cmd     = &cobra.Command{
		Use: "telnet host port",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return errors.New("host and port params required")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			client, err := client.NewClient(&client.ConnectOptions{
				Host:    args[0],
				Port:    args[1],
				Timeout: time.Duration(timeout) * time.Second,
			})
			if err != nil {
				log.Fatalf("Error: %s\n", err)
			}
			defer client.Close()

			client.Run()
		},
	}
)

func init() {
	cmd.PersistentFlags().IntVar(&timeout, "timeout", 10, "Connection timeout in seconds")
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
