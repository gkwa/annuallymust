package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
	"github.com/gkwa/annuallymust/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of annuallymust",
	Long:  `All software has versions. This is annuallymust's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
