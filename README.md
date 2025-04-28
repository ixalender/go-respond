# go-respond

[![Go Reference](https://pkg.go.dev/badge/github.com/ixalender/go-respond.svg)](https://pkg.go.dev/github.com/ixalender/go-respond)
[![GitHub release](https://img.shields.io/github/v/release/ixalender/go-respond)](https://github.com/ixalender/go-respond/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ixalender/go-respond/blob/main/LICENSE)

Lightweight HTTP response helper for Go REST APIs with zero dependencies.

## Installation

```bash
go get github.com/ixalender/go-respond
```

## Quick Examples

### Basic Usage

```go
import "github.com/ixalender/go-respond"

func GetUser(w http.ResponseWriter, r *http.Request) {
    user := User{ID: 123, Name: "John"}
    respond.Okay(w, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    // ...
    respond.Respond(w, newUser, http.StatusCreated)
}

func UserNotFound(w http.ResponseWriter) {
    respond.NotFound(w, "User not found")
}
```

### Available Methods

#### Success Responses

```go
respond.OK(w)                                   // 200 OK (empty)
respond.Okay(w, data)                           // 200 OK with data
respond.Accepted(w, data)                       // 202 Accepted
respond.Respond(w, data, http.StatusCreated)    // Any status code
```

#### Error Responses

```go
respond.Error(w, 400, "Bad request")  // Custom error
respond.BadRequest(w)                 // 400
respond.InternalError(w)              // 500
```

### Response Format

All error responses follow this JSON structure:

```json
{
  "message": "Error description",
  "status": 404
}
```

## Features

✔ Zero dependencies  
✔ Proper HTTP headers and status codes  
✔ Auto JSON encoding with HTML escaping  
✔ Consistent error response format

## License

MIT
