
# Distributed Rate Limiter (Go + Redis)

A **distributed token-bucket rate limiter** implemented in Go, using **Redis** and **Lua scripting** to enforce API request limits consistently across stateless services. 

---

##  Features

* Token Bucket rate-limiting algorithm
* Distributed state using Redis (works across multiple servers)
* Atomic request validation using Redis Lua scripts
* Stateless Go HTTP middleware
* Simple HTTP API for testing and demonstration

---

##  How It Works

* Each client is identified by a key (currently client IP).
* Redis stores a **token bucket** per client:

  * remaining tokens
  * last refill timestamp
* A Lua script:

  * refills tokens based on elapsed time
  * caps tokens at bucket capacity
  * consumes a token if available
* Lua execution in Redis guarantees **atomicity**, even under concurrent requests.

---

##  Architecture Overview

```
Client
  │
  ▼
Go HTTP Server (Middleware)
  │
  ▼
Redis (Token Bucket State)
  │
  ▼
Lua Script (Atomic Logic)
```

* Go servers remain **stateless**
* Redis acts as the shared coordination layer

---

## 📁 Project Structure

```
Rate_Limitor/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── internal/
│   ├── middleware/
│   │   ├── ratelimiter.go
│   │   └── token_bucket.lua
│   └── redis/
│       └── client.go
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

## Configuration

Configuration is loaded **from environment variables**.

### `.env.example :` 

```env
REDIS_ADDRESS=your_redis_host:port
REDIS_USERNAME=default
REDIS_PASSWORD=your_password

BUCKET_CAPACITY=10
REFILL_RATE=1
```

---

##  Running the Project

From the project root:

```bash
go run cmd/server/main.go
```
Expected output:

```
Connected to Redis
Server running on :8080
```
#ScreenShot :
<img width="1284" height="166" alt="image" src="https://github.com/user-attachments/assets/9feb8fb0-895b-466a-af45-121a6820eaaa" />


---

##  Testing the Rate Limiter

### Single request

```
curl.exe http://localhost:8080/api
```

Response:

```
Request allowed
```
#ScreenShot:
<img width="1001" height="220" alt="image" src="https://github.com/user-attachments/assets/8925cf1c-0331-49f4-81ef-122e21ed0cab" />


---

### Exceeding the limit

```
for ($i = 1; $i -le 20; $i++) { curl.exe http://localhost:8080/api }
```

After tokens are exhausted:

```
429 Too Many Requests
```
#ScreenShot:

<img width="1392" height="603" alt="image" src="https://github.com/user-attachments/assets/45eb454b-7f5c-4eea-bb6c-239b680d2c55" />


---

##  Inspecting Redis State 

Using `redis-cli`:

```redis
KEYS rate_limit:*
HGETALL rate_limit:<client_id>
```

You can observe:

* remaining tokens
* refill behavior over time

#ScreenShot :
<img width="1181" height="182" alt="image" src="https://github.com/user-attachments/assets/895b82f1-dd70-4361-aab9-1dfe9d4374b3" />


---

##  Design Decisions

* **Token Bucket**: allows short bursts while enforcing a steady rate
* **Redis + Lua**: ensures atomicity without locks
* **No background jobs**: token refill is computed lazily on request
* **Minimal API surface**: infrastructure-focused, not UI-driven


---

##  Use Cases

* API rate limiting
* Abuse prevention
* Traffic shaping
* Backend infrastructure learning project


