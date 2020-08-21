package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"kocli/pkg/auth"
	"os"
)

func init() {
	loginCmd.PersistentFlags().StringP("user", "u", "", "username")
	loginCmd.PersistentFlags().StringP("password", "p", "", "password")
	_ = loginCmd.MarkPersistentFlagRequired("user")
	_ = loginCmd.MarkPersistentFlagRequired("password")
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		err := auth.WithCredential(user, password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("login success")
		os.Exit(0)
	},
}
