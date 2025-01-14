# Sikkerbase: Database Backup Utility

**Sikkerbase** is a command-line interface (CLI) utility designed to simplify the process of backing up and restoring databases. It supports multiple database management systems (DBMS) like PostgreSQL, MySQL, and more. The tool provides features such as scheduling, compression, cloud storage integration, and detailed logging.

## Features

- **Database Connectivity**:
  - Support for multiple DBMS: PostgreSQL, MySQL.
  - Connection testing and validation.
- **Backup Operations**:
  - Supports full backups.
  - Compresses backup files to save space.
- **Storage Options**:
  - Local storage for backups.
  - Planned cloud integration (AWS S3, Google Cloud, Azure).
- **Logging and Notifications**:
  - Logs all backup activities with timestamps and statuses.
  - Notifications for success or failure (Slack integration planned).
- **Restore Operations**:
  - Restore from backup files.
  - Planned selective restore of specific tables.

## Installation

### Prerequisites

- Go 1.18 or higher
- PostgreSQL and/or MySQL client tools:
  - `pg_dump` for PostgreSQL
  - `mysqldump` for MySQL

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/gboliknow/sikkerbase.git
   cd sikkerbase
   ```
2. Build the CLI:
   ```bash
   go build -o sikkerbase
   ```
3. Ensure `pg_dump` and `mysqldump` are installed and accessible in your PATH.

## Usage

### Backup

To back up a PostgreSQL or MySQL database:

```bash
./sikkerbase backup \
  --dbtype=<postgres|mysql> \
  --host=<host> \
  --username=<username> \
  --password=<password> \
  --dbname=<database_name> \
  --output=<output_directory>
```

#### Example:

- **PostgreSQL**:
  ```bash
  ./sikkerbase backup --dbtype=postgres --host=localhost --username=admin \
    --password=secret --dbname=mydb --output=./backups
  ```

- **MySQL**:
  ```bash
  ./sikkerbase backup --dbtype=mysql --host=localhost --username=root \
    --password=secret --dbname=mydb --output=./backups
  ```

### Restore (Planned Feature)
Restore operations will allow you to recover databases from previously created backups. Selective restoration will be supported in future releases.

## Project Structure

```
.
├── cmd
│   └── sikkerbase
│       ├── main.go          # CLI entry point
├── db
│   ├── postgres.go         # PostgreSQL backup functions
│   ├── mysql.go            # MySQL backup functions
├── internal
│   ├── logger              # Logging utilities
│   └── storage             # Storage handling (local/cloud)
├── scripts
│   └── setup.sh            # Helper scripts
├── README.md               # Project documentation
```

## Dependencies

- **Cobra**: CLI framework for creating commands.
- **os/exec**: For executing external commands (`pg_dump` and `mysqldump`).

## Roadmap

- Add incremental and differential backup support.
- Integrate cloud storage (AWS S3, Google Cloud Storage).
- Implement selective restore.
- Add Slack/Email notifications.

## Contributing

Contributions are welcome! Feel free to submit pull requests or raise issues in the repository.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

**Happy Backing Up!**

