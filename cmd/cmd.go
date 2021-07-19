package cmd

import (
	"fmt"
	"lgmontenegro/crypto_bot/internal/application"
	"lgmontenegro/crypto_bot/internal/config"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg     config.Config
	app     application.Application
	env     string
	rootCmd = &cobra.Command{
		Use:   "crypto_bot",
		Short: "A bot for crypto pair value crawler",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			app.Bootstrap(cfg)
			app.Start()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configFilePath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(path.Dir(configFilePath))

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
		os.Exit(1)
	}
}
