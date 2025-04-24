# ELO Sport Competition Platform

A RESTful API for managing athlete competitions and score tracking across various martial arts styles, using the ELO rating system for fair matchmaking.

## Overview

This platform allows athletes to:
- Register in multiple martial arts disciplines
- Challenge other athletes to bouts
- Track their ELO rating progression over time
- Follow other athletes
- Participate in competitions across various gyms

The system uses the ELO rating system (similar to chess ratings) to provide fair matchmaking and track progression in each discipline separately.

## Features

- **Athlete Management**: Register, update, and manage athlete profiles
- **Style Management**: Support for multiple martial arts styles
- **Bout System**: Challenge, accept, and record outcomes of bouts
- **ELO Rating**: Automatic calculation and history tracking of athlete ratings
- **Social Features**: Follow other athletes and view activity feeds
- **Gym Management**: Register and associate with training facilities

## Technology Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Libraries**:
  - gorilla/mux - HTTP router
  - jmoiron/sqlx - Enhanced database operations
  - PostgreSQL driver

## Project Structure

```
elo-sport-comp/
├── interfaces/      # Interface definitions for services
├── models/          # Data models and structs
├── repositories/    # Database interaction layer
├── router/          # HTTP routing configuration
├── services/        # Business logic and handlers
├── utils/           # Shared utilities and helpers
└── databaseScripts/ # Database setup, DDL, and data insertion for testing
```

## Prerequisites

- Go 1.16+
- PostgreSQL 12+
- Git

## Installation and Setup

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/elo-sport-comp.git
cd elo-sport-comp
```

### 2. Set up the database

Run the database setup script:

```bash
cd databaseScripts
chmod +x setup_database.sh
./setup_database.sh
```

This script will:
- Create the necessary PostgreSQL database
- Apply the schema from CreateDBScript.sql
- Insert test data (optional) from InsertTestData.sql

### 3. Configure the application

Create a `.env` file in the root directory with the following variables (adjust as needed):

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=elo_sport_comp
SERVER_PORT=8080
```

### 4. Build and run the application

From the project root:

```bash
go build -o elo-sport-comp
./elo-sport-comp
```

The API will be available at http://localhost:8080/api/v1

## API Endpoints

### Athletes

- `GET /api/v1/athletes` - Get all athletes
- `GET /api/v1/athlete/{athlete_id}` - Get a specific athlete
- `POST /api/v1/athlete` - Create a new athlete
- `PUT /api/v1/athlete/{athlete_id}` - Update an athlete
- `DELETE /api/v1/athlete/{athlete_id}` - Delete an athlete
- `GET /api/v1/athlete/{athlete_id}/record` - Get athlete's record
- `POST /api/v1/athlete/authorize` - Authenticate an athlete
- `POST /api/v1/athletes/follow` - Follow an athlete
- `DELETE /api/v1/athletes/{followerId}/{followedId}/unfollow` - Unfollow an athlete
- `GET /api/v1/athletes/following/{id}` - Get followed athletes

### Bouts

- `GET /api/v1/bouts` - Get all bouts
- `GET /api/v1/bout/{bout_id}` - Get a specific bout
- `POST /api/v1/bout` - Create a new bout challenge
- `PUT /api/v1/bout/{bout_id}` - Update a bout
- `DELETE /api/v1/bout/{bout_id}` - Delete a bout
- `PUT /api/v1/bout/{bout_id}/accept` - Accept a bout challenge
- `PUT /api/v1/bout/{bout_id}/decline` - Decline a bout challenge
- `PUT /api/v1/bout/{bout_id}/complete/{referee_id}` - Complete a bout
- `PUT /api/v1/bout/cancel/{bout_id}/{challenger_id}` - Cancel a bout
- `GET /api/v1/bouts/pending/{athlete_id}` - Get athlete's pending bouts
- `GET /api/v1/bouts/incomplete/{athlete_id}` - Get athlete's incomplete bouts

### Outcomes

- `GET /api/v1/outcomes` - Get all outcomes
- `GET /api/v1/outcome/{outcome_id}` - Get a specific outcome
- `POST /api/v1/outcome` - Create an outcome
- `GET /api/v1/outcome/bout/{bout_id}` - Get outcome for a bout
- `POST /api/v1/outcome/bout/{bout_id}` - Create outcome for a bout

### Styles

- `GET /api/v1/styles` - Get all martial art styles
- `POST /api/v1/style` - Create a new style
- `POST /api/v1/style/athlete/{athlete_id}` - Register athlete to a style
- `POST /api/v1/styles/athlete/{athlete_id}` - Register athlete to multiple styles
- `GET /api/v1/styles/common/{athlete_id}/{challenger_id}` - Get common styles between athletes

### Athlete Scores

- `GET /api/v1/score/{athlete_id}` - Get all scores for an athlete
- `GET /api/v1/score/{athlete_id}/all` - Get all scores for an athlete
- `GET /api/v1/score/{athlete_id}/style/{style_id}` - Get athlete's score for a specific style
- `GET /api/v1/score/{athlete_id}/style/{style_id}/history` - Get historical scores by style and athlete

### Feed

- `GET /api/v1/feed/{athlete_id}` - Get activity feed for an athlete

### Gyms

- `GET /api/v1/gyms` - Get all gyms
- `GET /api/v1/gym/{gym_id}` - Get a specific gym
- `POST /api/v1/gym` - Create a new gym

## License

[MIT License](LICENSE)