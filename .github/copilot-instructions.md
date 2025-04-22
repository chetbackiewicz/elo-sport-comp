## Directory Structure
```
elo-sport-comp/
├── interfaces/      # Interface definitions for services
├── models/          # Data models and structs
├── repositories/    # Database interaction layer
├── router/          # HTTP routing configuration
├── services/        # Business logic and handlers
└── utils/           # Shared utilities and helpers
└── databaseScripts/ # Database setup, DDL, and data insertion for testing
```
## Coding Requirements

- All routes should go in the router.go file
- When creating new models, ensure to reference the CreateDBScript.sql DDL to understand the fields returned from the database
- All endpoints under `/api/v1` with RESTful naming
- Use `sqlx` for database operations
- Use prepared statements for SQL
- Use COALESCE for NULL handling in SQL queries

## When Answering
- Don't put comments above suggested terminal commands