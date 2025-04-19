# Elo Sport Competition Style Guide and Architecture

## Project Overview
Elo Sport Competition is a Go-based web application that manages athletic competitions, bout tracking, and athlete rankings. The application follows a clean architecture pattern with clear separation of concerns.

## Architecture

### Directory Structure
```
elo-sport-comp/
├── interfaces/     # Interface definitions for services
├── models/         # Data models and structs
├── repositories/   # Database interaction layer
├── router/         # HTTP routing configuration
├── services/       # Business logic and handlers
└── utils/         # Shared utilities and helpers
```

### Layer Responsibilities

1. **Interfaces Layer** (`interfaces/`)
   - Defines service contracts
   - Ensures loose coupling between layers
   - Example: `AthleteService`, `BoutService`

2. **Models Layer** (`models/`)
   - Defines data structures
   - Uses struct tags for JSON serialization and DB mapping
   - Contains no business logic
   - Example: `Athlete`, `Bout`, `Record`

3. **Repository Layer** (`repositories/`)
   - Handles database operations
   - Uses `sqlx` for database interactions
   - Implements CRUD operations
   - Returns domain models
   - Example: `AthleteRepository`, `BoutRepository`

4. **Service Layer** (`services/`)
   - Implements business logic
   - Handles data validation
   - Coordinates between repositories
   - Example: `athleteService`, `boutService`

5. **Handler Layer** (`services/`)
   - HTTP request handling
   - Request parsing and response formatting
   - Uses service layer for business logic
   - Example: `AthleteHandler`, `BoutHandler`

6. **Router Layer** (`router/`)
   - Defines API routes
   - Configures middleware
   - Groups related endpoints
   - Uses `/api/v1` prefix for all routes

## Database Practices

### PostgreSQL Configuration
- Uses `sqlx` for enhanced database operations
- Connection pooling with configured limits:
  - Max open connections: 5
  - Max idle connections: 5
  - Connection max lifetime: 5 minutes
- Environment-based configuration via `.env` file

### Query Patterns
- Use prepared statements for safety
- Implement proper error handling
- Use transactions where necessary
- Prefer explicit column naming
- Use COALESCE for NULL handling

## API Design

### Endpoint Structure
- RESTful endpoints under `/api/v1`
- Resource-based naming (e.g., `/athletes`, `/bouts`)
- Consistent HTTP method usage:
  - GET: Retrieval
  - POST: Creation
  - PUT: Updates
  - DELETE: Deletion

### Response Format
```json
// Success Response
{
  "data": {},
  "message": "Success message"
}

// Error Response
{
  "error": "Error message",
  "status": 400
}
```

### Error Handling
- Use appropriate HTTP status codes
- Return descriptive error messages
- Log errors with sufficient context
- Handle both application and database errors

## Code Style

### Naming Conventions
- PascalCase for exported types and functions
- camelCase for internal types and functions
- Use descriptive, action-based names for functions
- Prefix interfaces with 'I' (e.g., `IAthleteService`)

### Function Design
- Single responsibility principle
- Clear input validation
- Proper error handling and propagation
- Consistent return types
- Document complex logic

### Comments and Documentation
- Package-level documentation
- Interface documentation
- Complex function documentation
- Business logic explanation where necessary

## Testing

### Test Organization
- Unit tests alongside source files
- Integration tests in separate directory
- Use table-driven tests
- Mock interfaces for testing

### Test Naming
- Format: `Test<Function>_<Scenario>`
- Clear description of test case
- Group related tests

## Middleware

### Standard Middleware
- Logging middleware for request/response
- CORS configuration
- Authentication/Authorization
- Request ID tracking
- Response headers standardization

## Security Practices

### Authentication
- Proper password hashing
- Session management
- Token-based authentication
- Secure cookie handling

### Data Protection
- Input validation
- SQL injection prevention
- XSS protection
- CSRF protection

## Performance Considerations

### Database Optimization
- Proper indexing
- Query optimization
- Connection pooling
- Prepared statements

### Response Optimization
- Proper error handling
- Efficient JSON serialization
- Response caching where appropriate
- Pagination for large datasets

## Development Workflow

### Version Control
- Feature branches
- Meaningful commit messages
- Pull request reviews
- Version tagging

### Environment Configuration
- Use `.env` files for configuration
- Separate configurations for different environments
- Never commit sensitive data
- Document all environment variables

## Dependencies

### Core Dependencies
- `github.com/gorilla/mux`: Router
- `github.com/jmoiron/sqlx`: Database operations
- `github.com/lib/pq`: PostgreSQL driver
- `github.com/joho/godotenv`: Environment configuration

### Dependency Management
- Use Go modules
- Pin dependency versions
- Regular security updates
- Document major dependency changes

## Best Practices

### Code Organization
1. Group related functionality
2. Maintain clear layer separation
3. Use interfaces for flexibility
4. Keep functions focused and small

### Error Handling
1. Use custom error types
2. Proper error wrapping
3. Consistent error messages
4. Appropriate error logging

### Configuration
1. Environment-based configuration
2. Secure credential management
3. Configuration validation
4. Fallback values where appropriate

### Logging
1. Consistent log levels
2. Structured logging
3. Contextual information
4. Performance impact consideration

## AI Development Guidelines

When extending this project, AI assistants should:

1. **Maintain Architecture**
   - Follow existing layer separation
   - Use established patterns
   - Maintain interface contracts

2. **Code Generation**
   - Follow naming conventions
   - Include proper documentation
   - Implement error handling
   - Add appropriate tests

3. **Database Operations**
   - Use prepared statements
   - Implement proper transactions
   - Follow existing query patterns
   - Consider performance impact

4. **API Extensions**
   - Follow REST principles
   - Maintain URL structure
   - Implement proper validation

5. **Security Considerations**
   - Validate all inputs
   - Protect sensitive data
   - Follow authentication patterns
   - Implement proper authorization
