-- KEYS[1]  = rate limit key (e.g. rate_limit:<client_id>)
-- ARGV[1]  = bucket capacity (max tokens)
-- ARGV[2]  = refill rate (tokens per second)
-- ARGV[3]  = current timestamp (unix seconds)

local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

-- Read existing bucket state
local data = redis.call("HMGET", key, "tokens", "last_refill")
local tokens = tonumber(data[1])
local last_refill = tonumber(data[2])

-- Initialize bucket for first-time client
if tokens == nil then
    tokens = capacity
    last_refill = now
end

-- Refill tokens based on elapsed time
local elapsed = now - last_refill
local refill = elapsed * refill_rate
tokens = math.min(capacity, tokens + refill)

-- Reject request if no tokens left
if tokens < 1 then
    return 0
end

-- Consume one token
tokens = tokens - 1

-- Save updated state
redis.call("HMSET", key,
    "tokens", tokens,
    "last_refill", now
)

-- Allow request
return 1


