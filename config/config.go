package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

var (
    // Redis connection
    RedisAddress  string
    RedisUsername string
    RedisPassword string

    // Rate limiter configuration
    BucketCapacity      int
    RefillRatePerSecond int

    // Key prefix for rate limiting
    RateLimitKeyPrefix = "rate_limit:"
)

func init() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Warning: .env file not found")
    }

    // Load Redis config
    RedisAddress = os.Getenv("REDIS_ADDRESS")
    RedisUsername = os.Getenv("REDIS_USERNAME")
    RedisPassword = os.Getenv("REDIS_PASSWORD")

    // Load rate limiter config
    BucketCapacity, _ = strconv.Atoi(getEnvOrDefault("BUCKET_CAPACITY", "10"))
    RefillRatePerSecond, _ = strconv.Atoi(getEnvOrDefault("REFILL_RATE", "1"))

    // Validate required fields
    if RedisAddress == "" {
        log.Fatal("REDIS_ADDRESS is required")
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}




