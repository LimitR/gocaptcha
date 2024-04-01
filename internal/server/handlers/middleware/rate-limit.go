package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type TTL interface {
	Put(k string)
	Get(k string) *int64
	MaxTtl() int
}

func RateLimit(ttl TTL) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ok := ttl.Get(c.IP())
		if ok != nil {
			c.Set("Retry-After", strconv.Itoa((ttl.MaxTtl()-int(time.Now().Unix()-*ok))*1000))
			return c.Status(429).Send([]byte{})
		}

		ttl.Put(c.IP())

		return c.Next()
	}
}
