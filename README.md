# SimplePay

A payment processing system with database migrations.

## Database Setup

### Prerequisites
- PostgreSQL database
- Go 1.23.3 or later

### Environment Variables
Set the `DATABASE_URL` environment variable to your PostgreSQL connection string:
```bash
export DATABASE_URL="postgres://username:password@localhost:5432/simplepay?sslmode=disable"
```

### Running Migrations

#### Option 1: Automatic Migration (Recommended)
When you start the application, migrations will run automatically:
```bash
go run cmd/api/main.go
```

#### Option 2: Manual Migration
Run migrations manually using the migration command:
```bash
go run cmd/migration/main.go -up
```

## Database Schema

The system creates the following tables:

1. **merchants** - Store merchant information
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `name` (VARCHAR(100))
   - `category` (VARCHAR(20))
   - `status` (VARCHAR(20))
   - `details` (JSON)

2. **keys** - Store merchant API keys
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `merchant_id` (VARCHAR(10), FOREIGN KEY)
   - `key` (VARCHAR(10))

3. **providers** - Payment providers
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `name` (VARCHAR(100))

4. **terminals** - Merchant terminals
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `merchant_id` (VARCHAR(10), FOREIGN KEY)
   - `provider_id` (VARCHAR(10), FOREIGN KEY)

5. **orders** - Payment orders
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `amount` (BIGINT)
   - `status` (VARCHAR(20))
   - `currency` (VARCHAR(3))
   - `order_details` (JSON)
   - `merchant_id` (VARCHAR(10), FOREIGN KEY)

6. **payments** - Payment transactions
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `order_id` (VARCHAR(10), FOREIGN KEY)
   - `merchant_id` (VARCHAR(10), FOREIGN KEY)
   - `amount` (BIGINT)
   - `currency` (VARCHAR(3))
   - `status` (VARCHAR(20))
   - `provider_id` (VARCHAR(10), FOREIGN KEY)
   - `forced_provider` (VARCHAR(10))
   - `terminal_id` (VARCHAR(10), FOREIGN KEY)

7. **provider_calls** - Provider API call logs
   - `id` (VARCHAR(10), PRIMARY KEY)
   - `payment_id` (VARCHAR(10), FOREIGN KEY)
   - `provider_request` (JSON)
   - `provider_response` (JSON)

## Migration System

The migration system uses Go files instead of SQL files. Each migration file contains:
- Both `Up` and `Down` functions in a single file
- Uses `tx.Exec()` to execute SQL queries
- Automatically tracks applied migrations in `schema_migrations` table

### Migration Files
- `001_create_merchants_table.go`
- `002_create_keys_table.go`
- `003_create_providers_table.go`
- `004_create_terminals_table.go`
- `005_create_orders_table.go`
- `006_create_payments_table.go`
- `007_create_provider_calls_table.go`

## Development

### Running the Application
```bash
go run cmd/api/main.go
```

### Running Tests
```bash
go test ./...
```

### Building
```bash
go build -o bin/api cmd/api/main.go
go build -o bin/migration cmd/migration/main.go
```
