# Golang Migrations

A SQL file-based migration runner for Go applications, built around a simple idea: **keep your migration scripts in Git while maintaining migration history and change detection in the database.**

I built this project after working on a data-driven .NET application that relied heavily on database objects such as tables, indexes, functions, procedures, and triggers. That project used `RoundhousE` for database migrations, and I liked its approach of treating database objects such as procedures and functions as versioned objects rather than requiring a new migration history file for every change.

I considered using `Goose` for this project, but wanted a migration system that could work naturally with an application's existing configuration and support different behavior for different types of database objects.

In particular, I wanted to:

- Use the application's existing configuration instead of configuring database credentials separately for the migration tool.
- Organize migrations into separate folders such as `tables`, `indexes`, `functions`, `procedures`, and `triggers`.
- Treat tables and indexes differently from objects such as functions and procedures.
- Avoid creating a new migration history file every time a function or procedure is modified.
- Detect when an already-applied migration has been modified.

So I built this as a **custom SQL file-based migration runner/template** that can be adapted to your own project.

## How Migrations Work

1. Migration files are organized under the configured migration folders, such as:

   ```text
   migrations/
   ├── tables/
   ├── indexes/
   ├── views/
   ├── functions/
   ├── procedures/
   └── triggers/
   ```

   If you prefer a different folder structure, you can modify `internal/migrations/runner.go`.

2. Every migration is recorded in the database along with a hash of the migration file.

3. The hash is used to determine whether an already-applied migration has been modified.

4. If `failOnModified` is enabled and an already-applied migration has changed, the migration process fails. This is useful for objects such as tables and indexes where you may want an explicit audit trail instead of silently modifying an existing migration.

   ```go
   tableFiles, err := getMigrationFiles("migrations/tables", true, histories)
   if err != nil {
   	return err
   }

   indexFiles, err := getMigrationFiles("migrations/indexes", true, histories)
   if err != nil {
   	return err
   }

   viewFiles, err := getMigrationFiles("migrations/views", false, histories)
   if err != nil {
   	return err
   }
   ```

   The migration folders and their behavior are fully configurable. You can keep separate folders for `tables`, `indexes`, `views`, `functions`, `procedures`, and `triggers`, use only the folders you need, or keep everything under a single `migrations` folder.

   You only need to change the `root` and `failOnModified` values in `internal/migrations/runner.go` to match your project's structure. Once configured, the rest of the migration logic remains the same.

5. For objects such as functions and procedures, migrations can be configured to run again when their SQL definition changes. This means you can maintain a single file for an object and update its definition without creating a new migration file for every change.

The goal is to keep the migration workflow **simple, Git-friendly, configurable, and suitable for applications that make heavy use of database objects**.

---

## Give it a TRY

### 1. Clone the repository in local

```bash
git clone https://github.com/Sahil2k07/golang-migration.git
```

### 2. Copy the example config file

```bash
cp configs/app.example.toml configs/app.toml
```

#### Edit these variables as per your local configurations

```toml
[app]
environment = "development"
json_logs = false
file_logging = false

[database]
db_host = "localhost"
db_port = "5432"
db_user = "postgres"
db_password = "shahil"
db_name = "migration"
```

#### Logging options:

- Set `json_logs = true` to enable JSON-formatted logging.
- Set `file_logging = true` to write logs to a file in addition to the terminal.

### 3. Download dependencies

```bash
go mod download
```

### 4. Apply the migrations

```bash
go run cmd/migration/main.go
```
