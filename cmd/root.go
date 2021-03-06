package cmd

import (
	"fmt"
	"os"

	"github.com/lbryio/sentinel/daemon"
	"github.com/lbryio/sentinel/env"
	"github.com/lbryio/sentinel/pools"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	//cobra.OnInitialize(config.InitializeConfiguration)
	rootCmd.PersistentFlags().BoolP("debugmode", "d", false, "turns on debug mode for the application command.")
	rootCmd.PersistentFlags().BoolP("tracemode", "t", false, "turns on trace mode for the application command, very verbose logging.")
	err := viper.BindPFlags(rootCmd.PersistentFlags())
	if err != nil {
		logrus.Panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "watcher",
	Short: "Watches over the blockchain",
	Long:  `Chain watch monitors all things related to the lbrycrd blockchain network`,
	Run: func(cmd *cobra.Command, args []string) {
		//Run the application
		config, err := env.NewWithEnvVars()
		if err != nil {
			logrus.Panic(err)
		}
		pools.CoinMineAPIKey = config.CoinMineAPIKey
		daemon.Start()
	},
}

// Execute executes the root command and is the entry point of the application from main.go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
