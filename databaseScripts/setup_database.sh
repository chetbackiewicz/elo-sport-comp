#!/bin/bash

# Stop PostgreSQL service if running
echo "Stopping PostgreSQL service..."
brew services stop postgresql@14

# Start PostgreSQL service
echo "Starting PostgreSQL service..."
brew services start postgresql@14

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
sleep 5

# Drop database if it exists and create a new one
echo "Dropping existing database if it exists..."
dropdb elo_sport --if-exists

echo "Creating new database..."
createdb elo_sport

# Run the database creation script
echo "Creating database schema..."
psql -d elo_sport -f CreateDBScript.sql

# Run the data insertion script
echo "Populating database with test data..."
psql -d elo_sport -f dataInserts/InsertTestData.sql

echo "Database setup complete!"

# Verify the setup by showing some basic stats
echo -e "\nDatabase Statistics:"
psql -d elo_sport << EOF
\echo '\nNumber of athletes:'
SELECT COUNT(*) FROM athlete;
\echo '\nNumber of bouts:'
SELECT COUNT(*) FROM bout;
\echo '\nNumber of completed bouts:'
SELECT COUNT(*) FROM bout WHERE completed = true;
\echo '\nNumber of pending bouts:'
SELECT COUNT(*) FROM bout WHERE completed = false;
\echo '\nAthletes and their scores:'
SELECT a.username, s.style_name, as2.score 
FROM athlete_score as2 
JOIN athlete a ON a.athlete_id = as2.athlete_id 
JOIN style s ON s.style_id = as2.style_id 
ORDER BY a.username, s.style_name;
EOF 