# Gophermart Backend API 🐹🛒

A production-ready robust backend service implemented in **Go (Golang)** reflecting Clean Architecture / Domain-Driven Design principles. Built to scale for e-commerce and inventory workflows. 

## 🏗️ Architecture & Core Features

- **Clean Architecture:** Organized strictly via Presentation (Handlers), Business Logic (Services), and Persistence (Repositories). Code remains highly testable and loosely coupled.
- **RESTful Endpoints:** Complete CRUD for Products (`GET`, `POST`, `PATCH`, `DELETE`).
- **Resilient Infrastructure:** Configured with `net/http` **Graceful Shutdown** via `os.Signal` catching to prevent data corruption during container rollouts.
- **Advanced Pagination & Soft Delete:** Repositories are equipped with efficient pagination limits (`LIMIT/OFFSET`) and "Soft Delete" logic ensuring relational history integrity (`deleted_at`).
- **Idiomatic Go DI:** Pure explicit Manual Dependency Injection—no reflect-based magic frameworks, ensuring max compilation-time safety.
- **Environment Driven:** Secrets and connections parsed securely via `godotenv`.

## 📂 Project Structure (Package-By-Feature)

```text
gophermart/
├── cmd/
│   └── api/                  # Application entry points
│       ├── main.go           # Bootstrapping, DI wiring, and Graceful Shutdown
│       └── router.go         # API layout and Chi router registration
├── db/
│   └── migrations/           # Raw .sql scripts for Schema modifications
├── internal/
│   ├── api/                  # Global unified response formatters
│   ├── config/               # .env loader and environment structures
│   ├── infrastructure/       # Database Connection drivers (PostgreSQL)
│   ├── product/              # The "Product" Domain Feature
│   │   ├── entity/           # Structs, validation tags, Domain Models (Create/Update rules)
│   │   ├── handler/          # HTTP Request parsing & mapping (Presentation)
│   │   ├── repository/       # PostgreSQL queries and raw data interaction
│   │   └── service/          # Core Business logic
│   └── utils/                # DRY utility functions (e.g., param parsing)
├── .env.example              # Example environments
└── Makefile                  # Development tooling commands
```

## 🚀 Quick Start Guide

### 1. Prerequisites
- **Go** >= 1.20
- **PostgreSQL** >= 14
- `golang-migrate` command-line tool

### 2. Configuration
Copy the `.env.example` file (or create `.env`) at the root of the project:
```bash
DB_CONN_STR=postgres://postgres:password@localhost:5432/gophermart?sslmode=disable
PORT=8080
```

### 3. Database Migrations
Run the schema setup utilizing the Makefile:
```bash
make migrate-up
```

### 4. Running the Server
You can kickstart the development environment using:
```bash
make run
```
*The server will spin up on `http://localhost:8080`!*

### 🛠️ Developer Tooling
Check the `Makefile` for available workflows:
- `make build` : Compiles the Go binary to `/bin/api`.
- `make test` : Runs unit testing recursively alongside coverages.
- `make migrate-down` : Reverts the most recent DB migration.
- `make migrate-create` : Scaffolds a new `.sql` file structure.
