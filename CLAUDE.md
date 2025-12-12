# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Run Commands

```bash
# Run the server
go run .

# Build the binary
go build -o event-api .

# Run with hot reload (if air is installed)
air
```

The server runs on `localhost:8080`.

## Architecture Overview

This is a Go REST API for event management built with the Gin framework and SQLite.

### Package Structure

- `main.go` - Entry point: initializes database and starts Gin server
- `db/` - Database initialization and table creation (SQLite via go-sqlite3)
- `models/` - Data models with database operations (Event, User)
- `routes/` - HTTP handlers and route registration
- `middlewares/` - JWT authentication middleware
- `utils/` - JWT token generation/verification and password hashing (bcrypt)

### Authentication Flow

1. User signs up via `POST /signup` (password hashed with bcrypt)
2. User logs in via `POST /login` (returns JWT token)
3. Protected routes require `Authorization` header with JWT token
4. Middleware extracts `userId` from token and sets it in Gin context

### API Endpoints

**Public:**
- `GET /events` - List all events
- `GET /events/:id` - Get single event
- `POST /signup` - Register user
- `POST /login` - Authenticate user

**Protected (require JWT):**
- `POST /events` - Create event
- `PUT /events/:id` - Update event (owner only)
- `DELETE /events/:id` - Delete event (owner only)
- `POST /events/:id/register` - Register for event
- `DELETE /events/:id/register` - Cancel registration

### Database Schema

Three tables in `api.db`:
- `users` (id, email, password)
- `events` (id, name, description, location, dateTime, user_id)
- `registrations` (id, event_id, user_id)

### Testing API

HTTP test files are in `api-test/` directory for use with REST client extensions.
