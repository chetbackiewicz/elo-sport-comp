package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() *sqlx.DB {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database configuration
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Debug: Print all environment variables
	log.Printf("Environment variables loaded:")
	log.Printf("  DB_HOST: %s", dbHost)
	log.Printf("  DB_PORT: %s", dbPort)
	log.Printf("  DB_USER: %s", dbUser)
	log.Printf("  DB_NAME: %s", dbName)
	if dbPassword == "" {
		log.Printf("  DB_PASSWORD: (empty)")
	} else {
		log.Printf("  DB_PASSWORD: (set)")
	}

	// Validate required environment variables
	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		log.Fatal("Missing required database configuration in .env file")
	}

	// Convert port to integer
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("Invalid DB_PORT in .env file: %v", err)
	}

	// Construct connection string with explicit database name and additional parameters
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, port, dbName)

	// Log the exact connection string being used (without password)
	log.Printf("Using connection string: postgres://%s:****@%s:%d/%s?sslmode=disable",
		dbUser, dbHost, port, dbName)

	// Open database connection
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * 60) // 5 minutes

	// Test the connection with retry
	var lastErr error
	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		lastErr = err
		log.Printf("Connection attempt %d failed: %v", i+1, err)
	}

	if lastErr != nil {
		log.Fatalf("Failed to connect to database after 3 attempts: %v\nPlease ensure:\n1. PostgreSQL is running\n2. The database '%s' exists\n3. The user '%s' has access to the database",
			lastErr, dbName, dbUser)
	}

	log.Println("Successfully connected to database")
	return db
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
