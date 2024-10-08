services:
    redis:
      container_name: "redis"
      image: redis:alpine
      command: redis-server /usr/local/etc/redis/redis.conf --requirepass ${REDIS_PASSWORD}
      ports:
        - "6379:6379"
    web:
      container_name: "redisapi"
      build:
        context: .
      ports:
        - "8080:8080"