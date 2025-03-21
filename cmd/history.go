/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/temryakov/pomodoro/app"
	"github.com/temryakov/pomodoro/constants"
	"github.com/temryakov/pomodoro/repository"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   constants.HistoryUse,
	Short: constants.HistoryShort,
	Long:  constants.HistoryLong,
	Run:   RunHistory,
}

func RunHistory(cmd *cobra.Command, args []string) {
	r := repository.NewRepository(app.InitDB())
	res, err := r.Get()
	if err != nil {
		panic(err)
	}
	app.GetHistoryList(res)
}

func init() {
	rootCmd.AddCommand(historyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// historyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// historyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
