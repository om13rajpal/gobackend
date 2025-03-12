# ⌬ Golang Backend Best Practices ⌬

This repository follows best practices for building a production-ready backend in Golang, implementing efficient and scalable patterns for API development.

## ▪ Features ▪

× RESTful API using Gin  
× Authentication with JWT  
× CRUD operations with MongoDB & PostgreSQL (pgx)  
× Input validation with Validator  
× Environment variable management using Godotenv  
× Email service using Gomail  
× Middleware for security and request validation  

## ▪ Tech Stack ▪

× Golang  
× Gin (Web Framework)  
× MongoDB & PostgreSQL (pgx)  
× JWT (Authentication)  
× Validator  
× Godotenv  
× Gomail  

## ▪ Installation ▪

1. Clone the repository:
   ```sh
   git clone https://github.com/om13rajpal/gobackend.git
   cd gobackend
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables:
   ```sh
   cp example.env .env
   ```
4. Run the application:
   ```sh
   go run cmd/server/main.go
   ```

## ▪ Environment Variables ▪
Create a `.env` file with the following:
```
PORT=3000
EMAIL=your-email@example.com
PASSWORD=your-password
MONGO_URI=mongodb://localhost:27017/
JWT_SECRET=golang
```

## ▪ Open-Source & License ▪
This project is open-source and licensed under the MIT License. Contributions are welcome!

━━━━━━━━━━━━━━━━━━━━━━━━━  
Crafted with precision ⌁ by **Om**
