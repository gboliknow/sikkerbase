package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sikkerbase",
	Short: "SikkerBase - A database backup utility",
	Long: `SikkerBase is a command-line tool to backup and restore databases securely,
supporting MySQL, PostgreSQL, MongoDB, and SQLite.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
