package cmd

import (
	"fmt"
	"lgmontenegro/crypto_bot/internal/application"
	"lgmontenegro/crypto_bot/internal/config"
	"os"
	"path"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg      config.Config
	app      application.Application
	Verbose  bool
	Pairs    []string
	Times    time.Duration
	URL      string
	EndPoint string

	rootCmd = &cobra.Command{
		Use:   "crypto_bot",
		Short: "A bot to watch crypto pair values",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CRYPTO_BOT")
			fmt.Println("You are running Crypto_bot_client v0.1.0!")
			fmt.Println("To stop it press Ctrl+C.")
			if !Verbose {
				fmt.Println("Don't forget to activate verbose mode if you want to watch the tickers")
			}

			app.Bootstrap(cfg)
			app.Start(Verbose)
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

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringSliceVarP(&Pairs, "pairs", "p", cfg.Pairs, "coins pairs")
	rootCmd.PersistentFlags().DurationVarP(&Times, "interval", "i", cfg.Times, "Time in seconds for each ticker recovery")
	rootCmd.PersistentFlags().StringVarP(&URL, "url", "u", cfg.URL, "URL to access the endpoint")
	rootCmd.PersistentFlags().StringVarP(&EndPoint, "endpoint", "e", cfg.Endpoint, "EndPoint address for the ticker")
	
	viper.BindPFlag("pairs", rootCmd.PersistentFlags().Lookup("pairs"))
	viper.BindPFlag("times", rootCmd.PersistentFlags().Lookup("times"))
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
	viper.BindPFlag("endpoint", rootCmd.PersistentFlags().Lookup("endpoint"))
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
