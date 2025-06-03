# Accounts Manager

Accounts Manager is a simple REST API written in Go for managing user accounts. It uses the Gin web framework and GORM with a SQLite database. Authentication is handled with JSON Web Tokens (JWT).

## Features

- Create, read, update and delete users
- JWT‑based authentication with access and refresh tokens
- SQLite database for storage

## Requirements

- Go 1.24+

## Getting Started

1. Install dependencies and run the application:

```bash
go run main.go
```

2. The server will start on `http://localhost:8080` and create a local `accounts.db` file if it does not exist.

## API

### Public Endpoints

- `POST /users` &ndash; create a new user.
- `POST /login` &ndash; authenticate and obtain an access and refresh token.

### Protected Endpoints

These require an `Authorization: Bearer <token>` header with a valid access token.

- `GET /users/:id` &ndash; retrieve a specific user.
- `PUT /users/:id` &ndash; update a user.
- `DELETE /users/:id` &ndash; remove a user.
- `GET /users/` &ndash; list all users.

### Refresh Token

`controllers/userController.go` contains a `RefreshToken` handler for issuing a new access token. A route for it is not currently configured but can be added as needed.

## Configuration

For simplicity the JWT signing key and database path are hard coded. In a production setting you should load configuration from environment variables or a configuration file.

## Next Steps

- Protect sensitive configuration values
- Add tests for handlers and middleware
- Document all endpoints in detail
