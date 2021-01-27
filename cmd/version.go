// Package cmd
// version Command
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/xwi88/version"
)

var (
	// versionFlag version Flag
	versionFlag *bool
)

var versionCMD = &cobra.Command{
	Use:     "version",
	Short:   "Print the executable file version",
	Long:    "",
	Example: "  version -v=false\n  version -v=true",
	Run: func(cmd *cobra.Command, args []string) {
		if *versionFlag {
			v := version.Get()
			log.Printf("[version] info:\n%v", v.StringWithIndent())
		}
	},
}
