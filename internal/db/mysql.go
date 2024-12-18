package db

import (
	"fmt"
	"os/exec"
	"time"
)


func BackupMySQL(host, user, password, dbName, output string) {
	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("%s/mysql_backup_%s.sql", output, timestamp)

	var cmd *exec.Cmd
	if dbName != "" {
		cmd = exec.Command("mysqldump", "-h", host, "-u", user, fmt.Sprintf("-p%s", password), dbName, "-r", backupFile)
	} else {
		cmd = exec.Command("mysqldump", "-h", host, "-u", user, fmt.Sprintf("-p%s", password), "--all-databases", "-r", backupFile)
	}

	fmt.Println("Starting MySQL backup...")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to backup MySQL:", err)
		return
	}
	fmt.Printf("MySQL backup successful! File saved at: %s\n", backupFile)
}
