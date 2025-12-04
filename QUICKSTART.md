# Quick Start Guide - ELO Sport Competition Platform

This guide will help you get the application running on your local machine in under 10 minutes.

## Prerequisites

Before you begin, ensure you have installed:
- **Go 1.19+** - [Download here](https://golang.org/dl/)
- **PostgreSQL 12+** - [Download here](https://www.postgresql.org/download/)
- **Git** - [Download here](https://git-scm.com/downloads)

Verify installations:
```bash
go version        # Should show Go 1.19 or higher
psql --version    # Should show PostgreSQL 12 or higher
git --version     # Should show Git version
```

---

## Step 1: Clone the Repository

```bash
git clone https://github.com/chetbackiewicz/elo-sport-comp.git
cd elo-sport-comp
```

---

## Step 2: Set Up the Database

### Option A: Automated Setup (macOS with Homebrew)

The project includes a setup script for macOS users:

```bash
cd databaseScripts
chmod +x setup_database.sh
./setup_database.sh
```

This script will:
- Start/restart PostgreSQL
- Create the `elo_sport` database
- Apply the schema
- Insert test data

### Option B: Manual Setup (All Platforms)

#### 2.1 Start PostgreSQL

**macOS (Homebrew):**
```bash
brew services start postgresql@14
```

**Linux (systemd):**
```bash
sudo systemctl start postgresql
```

**Windows:**
Start PostgreSQL from the Services panel or pgAdmin.

#### 2.2 Create the Database

```bash
# Create database (use psql or your preferred tool)
createdb elo_sport_comp

# Or using psql directly:
psql -U postgres -c "CREATE DATABASE elo_sport_comp;"
```

#### 2.3 Apply the Schema

```bash
cd databaseScripts
psql -U postgres -d elo_sport_comp -f CreateDBScript.sql
```

#### 2.4 (Optional) Load Test Data

```bash
psql -U postgres -d elo_sport_comp -f dataInserts/InsertTestData.sql
```

#### 2.5 Verify Database Setup

```bash
psql -U postgres -d elo_sport_comp -c "\dt"
```

You should see tables like `athlete`, `bout`, `outcome`, `style`, etc.

---

## Step 3: Configure Environment Variables

Create a `.env` file in the project root:

```bash
cd /path/to/elo-sport-comp
cp .env.example .env
```

Edit `.env` with your database credentials:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_actual_password
DB_NAME=elo_sport_comp

# Server Configuration
PORT=8000
```

**Note:** Replace `your_actual_password` with your PostgreSQL password.

---

## Step 4: Download Dependencies

```bash
go mod download
```

This will download all required Go packages.

---

## Step 5: Build the Application

```bash
go build -o elo-sport-comp
```

This creates an executable named `elo-sport-comp` (or `elo-sport-comp.exe` on Windows).

---

## Step 6: Run the Application

```bash
./elo-sport-comp
```

You should see output like:
```
2025/12/04 15:30:00 In Main App
2025/12/04 15:30:00 Successfully connected to database
2025/12/04 15:30:00 Server starting on port 8000...
```

---

## Step 7: Test the API

Open a new terminal and test the endpoints:

### Test 1: Get All Athletes
```bash
curl http://localhost:8000/api/v1/athletes
```

### Test 2: Get All Styles
```bash
curl http://localhost:8000/api/v1/styles
```

### Test 3: Get All Gyms
```bash
curl http://localhost:8000/api/v1/gyms
```

### Test 4: Create a New Athlete
```bash
curl -X POST http://localhost:8000/api/v1/athlete \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "username": "johndoe",
    "birth_date": "1990-01-15",
    "password": "securepass123",
    "email": "john.doe@example.com"
  }'
```

If you see JSON responses, congratulations! 🎉 The application is running successfully.

---

## Troubleshooting

### Issue: "Error loading .env file"
**Solution:** Make sure the `.env` file exists in the project root directory and is properly formatted.

### Issue: "Failed to connect to database"
**Solutions:**
1. Verify PostgreSQL is running:
   ```bash
   # macOS
   brew services list | grep postgresql
   
   # Linux
   sudo systemctl status postgresql
   ```

2. Check database credentials in `.env`
3. Verify database exists:
   ```bash
   psql -U postgres -l | grep elo_sport_comp
   ```

4. Test connection manually:
   ```bash
   psql -U postgres -d elo_sport_comp
   ```

### Issue: "dial tcp [::1]:5432: connect: connection refused"
**Solution:** PostgreSQL is not running. Start it using the commands in Step 2.1.

### Issue: Build errors
**Solution:** Ensure you have Go 1.19 or higher:
```bash
go version
go mod tidy
go build -v
```

### Issue: Port already in use
**Solution:** Either:
1. Stop the process using port 8000
2. Or change the `PORT` in `.env` to a different port (e.g., 8080)

---

## Using the API

### Base URL
All endpoints are prefixed with: `http://localhost:8000/api/v1`

### Available Endpoints

**Athletes:**
- `GET /athletes` - List all athletes
- `POST /athlete` - Create athlete
- `GET /athlete/{id}` - Get athlete details
- `PUT /athlete/{id}` - Update athlete
- `DELETE /athlete/{id}` - Delete athlete

**Bouts:**
- `GET /bouts` - List all bouts
- `POST /bout` - Create bout challenge
- `GET /bout/{id}` - Get bout details
- `PUT /bout/{id}/accept` - Accept challenge
- `PUT /bout/{id}/decline` - Decline challenge

**Styles:**
- `GET /styles` - List all martial arts styles
- `POST /style` - Create new style
- `POST /style/athlete/{athlete_id}` - Register athlete to style

**Gyms:**
- `GET /gyms` - List all gyms
- `POST /gym` - Create gym
- `GET /gym/{id}` - Get gym details

For complete API documentation, see [README.md](README.md) or [PROJECT_STATUS.md](PROJECT_STATUS.md).

---

## Development Mode

For development with auto-reload, you can use:

```bash
# Run directly without building
go run main.go

# Or use air for hot-reload (install first: go install github.com/cosmtrek/air@latest)
air
```

---

## Stopping the Application

Press `Ctrl+C` in the terminal where the application is running.

To stop PostgreSQL:

**macOS:**
```bash
brew services stop postgresql@14
```

**Linux:**
```bash
sudo systemctl stop postgresql
```

---

## Next Steps

1. ✅ Application is running
2. 📚 Read the [README.md](README.md) for detailed API documentation
3. 🔍 Check [PROJECT_STATUS.md](PROJECT_STATUS.md) for technical details
4. 🧪 Explore the API using Postman or curl
5. 🚀 Start building your martial arts competition platform!

---

## Need Help?

- Check the [PROJECT_STATUS.md](PROJECT_STATUS.md) for detailed technical information
- Review the API endpoints in [README.md](README.md)
- Check database schema in `databaseScripts/CreateDBScript.sql`
- Review sample data in `databaseScripts/dataInserts/`

---

**Last Updated:** December 4, 2025  
**Version:** 1.0  
**Tested On:** Go 1.19+, PostgreSQL 12+, macOS/Linux
