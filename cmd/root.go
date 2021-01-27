// Package cmd
// root command
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	confFile *string
)

// appCMD command app
var appCMD = &cobra.Command{
	Use:       "app",
	Short:     "go-tf service",
	Long:      "",
	Example:   "app version\n  app start",
	ValidArgs: []string{"start", "version"},
}

func init() {
	// add version cmd
	appCMD.AddCommand(versionCMD)
	versionFlag = versionCMD.Flags().BoolP("version", "v", true, "app version")

	// add start cmd
	appCMD.AddCommand(startCMD)
	confFile = startCMD.Flags().StringP("config", "c", "./profiles/dev/app.yml",
		"config file(default is ./profiles/dev/app.yml)")
}

// Execute execute cobra
func Execute() {
	if err := appCMD.Execute(); err != nil {
		panic(err.Error())
	}
}
