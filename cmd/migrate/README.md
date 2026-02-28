# Database Migrations

Direktori ini berisi file-file migration database menggunakan golang-migrate.

## Struktur File Migration

Setiap migration terdiri dari 2 file:
- `{version}_{name}.up.sql` - untuk menjalankan migration
- `{version}_{name}.down.sql` - untuk rollback migration

## Cara Menggunakan

### 1. Menjalankan Migration

```bash
# Menggunakan CLI tool
go run cmd/migrate/main.go up

# Atau install golang-migrate CLI
migrate -path database/migrations -database "postgresql://user:password@localhost:5432/dbname?sslmode=disable" up
```

### 2. Rollback Migration

```bash
# Rollback 1 step terakhir
go run cmd/migrate/main.go down
```

### 3. Cek Versi Migration

```bash
go run cmd/migrate/main.go version
```

### 4. Membuat Migration Baru

```bash
# Install golang-migrate CLI dulu
# MacOS
brew install golang-migrate

# Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Buat file migration baru
migrate create -ext sql -dir database/migrations -seq create_refresh_tokens_table
```

## Best Practices

1. **Jangan edit file migration yang sudah dijalankan** - Buat migration baru untuk perubahan
2. **Test migration up dan down** - Pastikan rollback berfungsi dengan baik
3. **Gunakan transaksi** - Wrap perubahan dalam BEGIN/COMMIT jika memungkinkan
4. **Naming convention** - Gunakan nama yang deskriptif: `create_table_name`, `add_column_to_table`, dll
5. **Backup database** - Selalu backup sebelum menjalankan migration di production

## Contoh Migration Baru

Jika ingin menambah kolom baru:

```sql
-- 000002_add_last_login_to_users.up.sql
ALTER TABLE users ADD COLUMN last_login TIMESTAMP WITH TIME ZONE;
```

```sql
-- 000002_add_last_login_to_users.down.sql
ALTER TABLE users DROP COLUMN last_login;
```