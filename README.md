# GoAPI
A RESTful API service built with Go using the Chi router for user management operations.

## Features
- User operations
- RESTful API design
- Lightweight and fast using Chi router
- Mock data layer for development
- Structured error handling
- Request/Response models

## Tech Stack

- **Go** 1.24.3

## Project Structure

```
goapi/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── app/
│   │   └── handlers/            # HTTP request handlers
│   │       ├── user_create.go
│   │       ├── user_delete.go
│   │       ├── user_get.go
│   │       ├── user_list.go
│   │       └── user_update.go
│   └── pkg/
│       ├── common/
│       │   ├── api/             # API request/response models
│       │   ├── constant/        # Constants (HTTP status codes)
│       │   └── errors/          # Error handling
│       ├── data/
│       │   ├── dao/             # Data access objects
│       │   └── mock/            # Mock data layer
│       └── models/
│           └── user.go          # User model
├── go.mod
└── README.md
```

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Users

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/v1/users/create` | Create a new user |
| POST | `/api/v1/users/list` | List all users |
| POST | `/api/v1/users/get` | Get a specific user |
| POST | `/api/v1/users/update` | Update a user |
| POST | `/api/v1/users/delete` | Delete a user |

### User Model

```json
{
  "id": "uuid-string",
  "name": "string",
  "email": "string",
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```