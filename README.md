# Laza

Laza is a marketplace for selling fashion products. This RESTful API built using Golang, PostgreSQL, Gin, and dependency injection as modularization.

## ⚡ Features

- Authentication & Authorization using JWT
- CRUD for all modules
- Manage Product for admin role
- Checkout & Payment for user role
- Get History order

## 💻 Built with

- [Golang](https://go.dev/): programming language
- [Gin](https://gin-gonic.com/): for handling HTTP requests and responses
- [JWT](https://github.com/golang-jwt/jwt) for authentication and authorization
- [Postgres](https://github.com/postgres/postgres) for DBMS

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/GoJav-Backend-Academy-B3/laza-backend.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Add Env

```sh
JWT_KEYS=

PSQL_USER= 
PSQL_PASS=  
PSQL_HOST=
PSQL_PORT=
PSQL_DBNAME=
PSQL_TIMEZONE=Asia/Jakarta
PSQL_TIMEOUT_1=1
PSQL_TIMEOUT_2=4
PSQL_TIMEOUT_3=7
# Use this if you have different schema target
PSQL_SEARCH_PATH=

# Setup for send mail
SMTP_HOST=
SMTP_PORT=
MAIL_USER=
MAIL_PASS=

# Setup for Cloudinary
CLOUD_NAME=
CLOUD_APIKEY=
CLOUD_SECRET=

APP_PORT=
BASE_URL=

```

9. Run Swagoo init asdkhgasdvoijsd
```
swag init -g cmd/server/router/router.go
```

4. Run the app

```bash
go run /cmd/server/main.go
```


🌟 You are all set!

