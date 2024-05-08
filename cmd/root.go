package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomodoro",
	Short: "🍅 Pomodoro CLI for your work and study",
	Long: `
🍅 Pomodoro CLI is a productivity cli application designed 
for your work and study. Stay focused and finish tasks effectively. 
Take more time to your life.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// var log *utils.Logger

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomodoro.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// log = utils.NewLogger()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
