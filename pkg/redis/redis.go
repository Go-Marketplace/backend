package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Redis struct {
	Client       *redis.Client

	connAttempts int
	connTimeout  time.Duration
}

func connectToRedis(ctx context.Context, conn string) (*redis.Client, error) {
	opts, err := redis.ParseURL(conn)
	if err != nil {
		log.Fatal("cannot parse redis url: ", err)
	}

	client := redis.NewClient(opts)
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return client, nil
}

func New(ctx context.Context, conn string) (*Redis, error) {
	r := &Redis{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	for r.connAttempts > 0 {
		client, err := connectToRedis(ctx, conn)
		if err == nil {
			r.Client = client
			break
		}

		log.Printf("Redis is trying to connect, attempts left: %d, err: %s", r.connAttempts, err)
		time.Sleep(r.connTimeout)
		r.connAttempts--
	}

	return r, nil
}

func (r *Redis) Close() {
	if r.Client != nil {
		r.Client.Close()
	}
}
