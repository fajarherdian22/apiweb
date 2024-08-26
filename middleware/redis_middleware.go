package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fajarherdian22/apiweb/db"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type responseWriter struct {
	gin.ResponseWriter
	body string
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body += string(b)
	return rw.ResponseWriter.Write(b)
}

var ctx = context.Background()

func RedisMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Request.URL.String()

		rdb := db.GetRedisClient()

		cachedResponse, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			fmt.Println("Cache not found")
			c.Writer = &responseWriter{c.Writer, ""}
			c.Next()
			body := c.Writer.(*responseWriter).body

			if c.Writer.Status() == http.StatusOK {
				rdb.Set(ctx, key, body, time.Minute*10)
				fmt.Println("Save into redis")
			}
			
		} else if err != nil {
			fmt.Println("Error connecting to Redis:", err)
			c.Next()
		} else {
			fmt.Println("Cache found")

			c.String(http.StatusOK, cachedResponse)
			c.Abort()
		}
	}
}
