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