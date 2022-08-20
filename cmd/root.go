/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/DE-Karpov/gitframe/core"
	"github.com/spf13/cobra"
)

var (
	exclude []string
	extensions []string
	format string
	include []string
	languages []string
	orderBy string
	repo string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitframe",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		core.Process(repo, orderBy, format, languages, extensions, include, exclude)
	 },
}

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
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.blameparser.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&repo, "repo", ".", `path to git repo`)
	rootCmd.PersistentFlags().StringVar(&orderBy,"order-by", "lines", `sort order, one of lines|commits|files`)
	rootCmd.PersistentFlags().StringVar(&format,"format", "tabular", `output format, one of tabular|json|json-lines|csv`)
	rootCmd.PersistentFlags().StringSliceVar(&exclude, "exclude", []string{}, "glop pattern to exclude files")
	rootCmd.PersistentFlags().StringSliceVar(&languages, "languages", []string{}, "calculate stats for given programming languages")
	rootCmd.PersistentFlags().StringSliceVar(&extensions, "extensions", []string{}, "calculate stats for given extensions")
	rootCmd.PersistentFlags().StringSliceVar(&include,"include", []string{}, "glop pattern to narrow down calculation")


}


