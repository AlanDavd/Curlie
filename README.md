# Curlie

A user-friendly HTTP API & server for generating curl commands.

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/alandavd/curlie.git
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the application:

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Usage

### Generate Curl Command

**Endpoint:** `POST /api/curl`

**Request Body:**

```json
{
  "method": "POST",
  "url": "https://api.example.com/data",
  "headers": {
    "Content-Type": "application/json",
    "Authorization": "Bearer token123"
  },
  "body": "{\"key\": \"value\"}",
  "queryParams": {
    "param1": "value1",
    "param2": "value2"
  }
}
```

**Response:**

```json
{
  "command": "curl -X POST -H 'Content-Type: application/json' -H 'Authorization: Bearer token123' -d '{\"key\": \"value\"}' 'https://api.example.com/data?param1=value1&param2=value2'"
}
```

## Project Structure

The project follows hexagonal architecture:

- `internal/domain`: Contains the core business logic and interfaces
- `internal/application`: Contains the service implementations
- `internal/infrastructure`: Contains the HTTP handlers and server setup
