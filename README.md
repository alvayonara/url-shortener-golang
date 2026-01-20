# Go URL Shortener

A simple URL shortener written in Go.  
This project was built to learn Go by building real backend service.  

---

## What this project does

- Create short URLs
- Redirect short codes to the original URL
- Store data in MySQL
- Keep the code easy to extend and reason about

---

## How it works

The application is split into three main layers:
- HTTP layer handles requests and responses
- Service layer contains business logic
- Store layer handles data persistence

---

## API

### Create short URL
POST /shorten  
Request body:

```json
{
  "url": "https://alvayonara.com"
}
```
Response:
```json
{
"code": "abc123",
"url": "https://alvayonara.com"
}
```

### Redirect
GET /{code}  
Response:
```json
302 Found
Location: https://alvayonara.com
```

---

## MySQL setup
Create database and table:
```json
CREATE DATABASE url_shortener;
USE url_shortener;

CREATE TABLE links (
    code VARCHAR(16) PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
Configure MySQL store in main.go:  
```json
store, err := store.NewMySQLStore(
"user:password@tcp(localhost:3306)/url_shortener",
)
```

## Run the server
```json
go run cmd/server/main.go
```
The server runs on:
```json
http://localhost:8080
```
