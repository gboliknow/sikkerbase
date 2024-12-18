package main

import "sikkerbase/cmd"

func main() {
	cmd.Execute()
}

// sikkerbase/
// │
// ├── cmd/                      # CLI commands
// │   ├── root.go               # Main CLI entry point
// │   ├── backup.go             # Backup command logic
// │   ├── restore.go            # Restore command logic
// │   └── testconn.go           # Test database connection command
// │
// ├── internal/                 # Core business logic
// │   ├── db/                   # Database handlers
// │   │   ├── mysql.go          # MySQL backup/restore logic
// │   │   ├── postgres.go       # PostgreSQL backup/restore logic
// │   │   ├── mongodb.go        # MongoDB backup/restore logic
// │   │   └── sqlite.go         # SQLite backup/restore logic
// │   │
// │   ├── storage/              # Storage logic
// │   │   ├── local.go          # Save backups locally
// │   │   ├── cloud.go          # Cloud upload/download logic
// │   │   └── aws.go            # AWS S3-specific implementation
// │   │
// │   ├── scheduler/            # Automatic scheduling
// │   │   └── cron.go           # Cron job setup for backups
// │   │
// │   └── notifications/        # Notifications (e.g., Slack)
// │       └── slack.go          # Slack notification logic
// │
// ├── pkg/                      # Shared utilities
// │   ├── compress.go           # Compression utilities (gzip/zip)
// │   ├── logger.go             # Logging utility
// │   └── config.go             # Configuration parser
// │
// ├── backups/                  # Default directory for storing backups
// │
// ├── main.go                   # Main entry point for CLI
// ├── go.mod                    # Go module configuration
// ├── go.sum                    # Go module dependencies
// └── README.md                 # Documentation
