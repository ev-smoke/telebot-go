/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/hirosassa/zerodriver"
	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TOKEN = os.Getenv("TOKEN")
)

// telebotGoCmd represents the telebotGo command
var telebotGoCmd = &cobra.Command{
	Use:   "telebotGo",
	Aliases: []string{"go"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contzains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := zerodriver.NewProductionLogger()

		fmt.Printf("telebotGo %s started", appVersion)
		telebotGo, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TOKEN,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			logger.Fatal().Str("Error", err.Error()).Msg("Please check TOKEN")
			return
		}
		telebotGo.Handle(telebot.OnText, func(m telebot.Context) error {
			logger.Info().Str("Payload", m.Text()).Msg(m.Message().Payload)
			payload := m.Text()

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm telebotGo %s!", appVersion))
				
			}
			return err
		})

		telebotGo.Start()
	},
}

func init() {
	rootCmd.AddCommand(telebotGoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// telebotGoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// telebotGoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
