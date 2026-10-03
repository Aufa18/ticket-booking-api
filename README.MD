# High-Concurrency Ticket Booking API 🎟️⚡

A highly concurrent RESTful API simulating a flash sale or "ticket war" system. Engineered to handle thousands of concurrent checkout requests without database crashes or overbooking.

## 🚀 Enterprise Architecture Highlights
*   **Redis Atomic Operations:** Utilizes Redis `DECR` for rapid, thread-safe quota deduction in-memory, achieving sub-millisecond response times.
*   **Goroutines & Channels:** Implements asynchronous background workers to process database writes, preventing the main API thread from blocking during traffic spikes.
*   **Preventing Race Conditions:** Safely handles concurrent ticket requests to ensure zero over-selling.
*   **Dockerized Infrastructure:** Uses Docker for reliable Redis containerization.

## 🛠️ Tech Stack
*   **Language:** Golang (Go 1.22)
*   **Framework:** Gin
*   **In-Memory Cache:** Redis (go-redis/v9)
*   **Database:** PostgreSQL (GORM)

## ⚙️ How to Run Locally

1. **Clone the repository:**
   ```bash
   git clone https://github.com/username/ticket-booking-api.git
   cd ticket-booking-api
2. **Start Redis via Docker:**
   ```Bash
   docker run --name redis-ticket -p 6379:6379 -d redis
3. **Setup Environment Variables:**
   Create a .env file:
   ```bash
   DB_HOST=localhost
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=ticket_db
   DB_PORT=5432
   REDIS_HOST=localhost:6379
4. **Run the Server:**
   ```Bash
   go mod tidy
   go run cmd/api/main.go