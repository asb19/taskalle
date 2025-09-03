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

---

### ⚡ Scalability

- **Horizantal Scaling**
    - each service can be scaled independently and load balanced using NGINX etc.
    - Later when we can use kubernetes and HPA can be applied for autoscaling each services
- **Database Scaling**
    - can add read replica for heavy traffic
    - can implement connection pooling for better connection utilization
    - sharding/partioning can be added later
- **Caching**
    - can add cache like Redis etc for storing frequently requested data

- **API Gateway**
    - can place an API gateway for all the services
    - it can help us with auth,rate limiting,request routing
    - ex: AWS API Gateway/Kong
- **gRPC**
    - for this assignment i have used REST as communication medium between the services
    - for higher throughput, w ecan use gRPC or message Queues.

---

## Enjoyed the Assignment.I could have made it more detailed or added more layers. But let me know the feedback. Thanks 😊.









