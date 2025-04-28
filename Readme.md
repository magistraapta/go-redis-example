# Query Caching with Redis

## Endpoints
I'm creating two endpoint to get user data from databse one with Redis and another one without Redis and compare their response time.

Without Redis
```bash
    GET localhost:8000/user/1
```

With Redis
```bash
    GET localhost:8000/redis/user/1
```

## Comparing the response time
![response-time](https://github.com/user-attachments/assets/d8be4b40-abdf-4fb5-a64a-701d574dc6cb)

## Tech stack
- Go: 1.24
- GORM
- Redis

## Installation

Run this code by write this command in Terminal.

```bash
    cd {project_name}
    gor run .
```