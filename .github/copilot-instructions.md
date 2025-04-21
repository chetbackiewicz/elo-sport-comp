## Directory Structure
```
elo-sport-comp/
├── interfaces/     # Interface definitions for services
├── models/         # Data models and structs
├── repositories/   # Database interaction layer
├── router/         # HTTP routing configuration
├── services/       # Business logic and handlers
└── utils/         # Shared utilities and helpers
```

## Key Development Guidelines

1. **Architecture**
   - Use clean architecture with strict layer separation (interfaces -> services -> repositories)
   - All database operations must use prepared statements and proper transactions
   - Services must implement their corresponding interfaces

2. **Code Style**
   - PascalCase for exported items, camelCase for internal
   - Every function must include error handling and logging
   - Use COALESCE for NULL handling in SQL queries
   - Keep functions focused and under 50 lines where possible

3. **API Design**
   - All endpoints under `/api/v1` with RESTful naming
   - Use appropriate HTTP methods (GET, POST, PUT, DELETE)
   - Implement pagination for list endpoints

4. **Database**
   - Use `sqlx` for database operations
   - Configure connection pools (max: 5 open, 5 idle)
   - Always use explicit column names in queries
   - Implement proper indexes for frequently queried fields

5. **Security**
   - Validate all input data
   - Use prepared statements for SQL
   - Implement proper authentication/authorization
   - Never log sensitive data

6. **Error Handling**
   - Return descriptive error messages
   - Log errors with context
   - Use appropriate HTTP status codes
   - Implement proper error wrapping

7. **Configuration**
   - Use `.env` for environment variables
   - Never commit sensitive data
   - Implement proper validation
   - Use secure credential management
