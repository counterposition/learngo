package main

import (
	"context"
	"log"

	"github.com/counterposition/learngo/internal/config"
	"github.com/counterposition/learngo/internal/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configuration config.Config

var rootCmd = &cobra.Command{
	Use:   "learn",
	Short: "Learning Go and its ecosystem",
	Run: func(cmd *cobra.Command, args []string) {
		client := db.NewEntClient(configuration.DatabaseUri)

		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			log.Panic(err)
		}
	},
}

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Panic(err)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
