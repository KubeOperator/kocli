package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ko",
		Short: "ko",
		Long:  `ko is command line tool for KubeOperator`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(hostCmd)
	rootCmd.AddCommand(clusterCmd)
	rootCmd.AddCommand(credentialCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(projectCmd)
}
