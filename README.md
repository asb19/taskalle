# Task & User Management Microservices

A simple microservices system built with **Go**, **PostgreSQL**, and **Gorilla Mux**.  
It demonstrates microservice design principles: separation of concerns, CRUD, pagination, filtering, REST communication, and Swagger documentation.

---

## 🚀 Features Implemented

- **Task Service**
  - CRUD operations on tasks (`/tasks`)
  - Pagination & filtering by status
  - `AssignedTo` field referencing User Service
  - Enriched task responses with `assigned_user` info
- **User Service**
  - GET endpoint for user details by id
- **Architecture**
  - Repository → Service → Handler layers
  - REST communication between Task and User services
  - Swagger documentation for API endpoints

---


---

## 🛠️ Setup & Run

### 1. Clone the repo

```bash
git clone https://github.com/asb19/taskalle.git
cd taskalle
docker compose up --build -d
```
## Services
  - Task Service → http://localhost:8080

  - User Service → http://localhost:8081

  - Postgres (taskdb) on port 5432

  - Postgres (userdb) on port 5433

### Migrations
    - Migrations run automatically on container startup.

## 📖 Swagger Documentation

Swagger docs are available at:

Task Service → http://localhost:8080/swagger/index.html








