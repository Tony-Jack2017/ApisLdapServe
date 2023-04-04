package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "apis",
	Short: "Apis is commandline tool for apis server.",
	Long: `A quick and convenient command line 
		   for running and compiling a series of 
		   actions for Apis Ldap backend services.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, errPrint := fmt.Fprintln(os.Stderr, err)
		if errPrint != nil {
			return
		}
		os.Exit(1)
	}
}
