# 🌐 API Design and Testing Best Practices

APIs (Application Programming Interfaces) are the backbone of modern software systems. Designing and testing APIs effectively ensures maintainability, scalability, and security.

---

## 1️⃣ API Design Principles

Good API design ensures **usability, maintainability, and consistency**.

### Key Principles:

1. **Simplicity and Clarity**  
   - Keep endpoints clear and meaningful.  
   - Example: `/users/{id}` instead of `/getUserById`.

2. **Consistency**  
   - Use consistent naming conventions and structure.  
   - HTTP methods:
     - `GET` → fetch data  
     - `POST` → create resource  
     - `PUT` → update resource  
     - `DELETE` → remove resource

3. **Versioning**  
   - Version APIs to manage changes without breaking existing clients.  
   - Example: `/v1/users`, `/v2/users`

4. **Error Handling**  
   - Return standard error codes and messages.  
   - Example:
     ```json
     {
       "error": "User not found",
       "code": 404
     }
     ```

5. **Idempotency**  
   - Repeated calls to the same API with the same input should produce the same result (especially for `PUT` and `DELETE`).

6. **Statelessness**  
   - APIs should not store client session data on the server. Each request contains all necessary information.

---

## 2️⃣ API Documentation

API documentation is essential for **developers to understand and consume APIs effectively**.

### Steps to Document an API:

1. **Describe Endpoints**  
   - Include method, path, parameters, request/response examples.
2. **Include Authentication Info**  
   - Explain how to authenticate (API keys, OAuth tokens, etc.).
3. **Provide Sample Requests and Responses**  
   - Example:
     ```bash
     curl -X GET "https://api.example.com/v1/users/1" \
          -H "Authorization: Bearer <token>"
     ```
     Response:
     ```json
     {
       "id": 1,
       "name": "John Doe",
       "email": "john@example.com"
     }
     ```
4. **Error Codes and Messages**  
   - Document all possible error responses for clarity.
5. **Tools for API Documentation**  
   - OpenAPI (Swagger)  
   - Postman Collections  
   - API Blueprint

---

## 3️⃣ API Security

APIs are vulnerable if not properly secured. Some best practices:

1. **Authentication and Authorization**  
   - Use OAuth 2.0, JWT, or API keys.  
   - Example: Bearer token in headers.

2. **Input Validation**  
   - Validate all inputs to prevent injection attacks.  
   - Example: Validate user-provided IDs, strings, and JSON structure.

3. **Rate Limiting**  
   - Protect APIs from abuse or DoS attacks.  
   - Example: Allow 100 requests per minute per user.

4. **HTTPS Only**  
   - Always serve APIs over HTTPS to ensure encryption.

5. **Error Handling**  
   - Avoid exposing sensitive server information in error messages.  
   - Example: Return generic error messages like `Internal Server Error` instead of stack traces.

6. **Logging and Monitoring**  
   - Track API usage, errors, and anomalies to detect issues early.

---

## 4️⃣ End-to-End Example: User Management API

### Step 1: Define Endpoints

| Method | Endpoint       | Description             |
|--------|----------------|-------------------------|
| GET    | /v1/users      | Get list of users       |
| GET    | /v1/users/{id} | Get a specific user     |
| POST   | /v1/users      | Create a new user       |
| PUT    | /v1/users/{id} | Update user information |
| DELETE | /v1/users/{id} | Delete a user           |

---

### Step 2: Implement API (Go + Gorilla Mux)

```go
package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: "1", Name: "John Doe", Email: "john@example.com"},
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/v1/users", getUsers).Methods("GET")
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
```
## Step 3: Test API with cURL
```
curl -X GET http://localhost:8080/v1/users
```

Expected Output:
```
[
  {
    "id": "1",
    "name": "John Doe",
    "email": "john@example.com"
  }
]
```
## Step 4: Secure API (JWT Example)

Use middleware to check JWT token before allowing access:
```
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token != "Bearer my-secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

```
Apply middleware to routes.

## Step 5: Document API

Provide examples, request/response formats, and authentication method in README or Swagger.

## Notes
| Aspect        | Best Practice                                           |
| ------------- | ------------------------------------------------------- |
| Design        | Consistent, simple, versioned, idempotent, stateless    |
| Documentation | Clear endpoints, sample requests/responses, error codes |
| Security      | Auth, input validation, HTTPS, rate limiting            |
| Testing       | Unit tests, integration tests, API contract tests       |

## Exercises 
- Extend the above API to include POST, PUT, and DELETE endpoints.

- Add input validation for email and name fields.

- Implement JWT authentication for all endpoints.

- Document all endpoints with request/response examples and error codes.

- Write unit tests for getUsers and POST /v1/users endpoints.
