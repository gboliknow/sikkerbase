package cmd

import (
	"fmt"
	"sikkerbase/internal/db"

	"github.com/spf13/cobra"
)

var dbType, host, username, password, dbName, output string

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting backup...")
		if dbType == "" || host == "" || username == "" || password == "" || dbName == "" || output == "" {
			fmt.Println("Error: All flags (dbtype, host, username, password, dbname, and output) must be provided.")
			return
		}
		switch dbType {
		case "mysql":
			db.BackupMySQL(host, username, password, dbName, output)
		case "postgres":
			db.BackupPostgres(host, username, password, dbName, output)
		default:
			fmt.Println("Unsupported database type")
		}
	},
}

func init() {
	backupCmd.Flags().StringVarP(&dbType, "dbtype", "t", "", "Type of database (mysql, postgres)")
	backupCmd.Flags().StringVarP(&host, "host", "H", "", "Database host")
	backupCmd.Flags().StringVarP(&username, "username", "u", "", "Database username")
	backupCmd.Flags().StringVarP(&password, "password", "p", "", "Database password")
	backupCmd.Flags().StringVarP(&dbName, "dbname", "d", "", "Name of the database to backup")
	backupCmd.Flags().StringVarP(&output, "output", "o", "", "Directory to save the backup file")

	backupCmd.MarkFlagRequired("dbtype")
	backupCmd.MarkFlagRequired("host")
	backupCmd.MarkFlagRequired("username")
	backupCmd.MarkFlagRequired("password")
	backupCmd.MarkFlagRequired("dbname")
	backupCmd.MarkFlagRequired("output")

	rootCmd.AddCommand(backupCmd)
}
