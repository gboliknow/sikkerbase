package db

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// BackupPostgres performs a PostgreSQL database backup using the pg_dump tool.
func BackupPostgres(host, user, password, dbName, output string) {
	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("%s/postgres_backup_%s.sql", output, timestamp)
	os.Setenv("PGPASSWORD", password)

	// Prepare the pg_dump command
	cmd := exec.Command("pg_dump",
		"-h", host,     
		"-U", user,     
		"-d", dbName,     
		"-F", "c",        
		"-f", backupFile,
	)

	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)

	fmt.Println("Starting PostgreSQL backup...")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error during PostgreSQL backup:", err)
		return
	}

	fmt.Printf("PostgreSQL backup successful! Backup file: %s\n", backupFile)
}
