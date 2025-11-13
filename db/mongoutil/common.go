package mongoutil

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	DefaultMaxPoolSize = 100
	DefaultMinRetry    = 3
)

func buildMongoURI(config *Config) string {
	credential := ""
	if config.Username != "" && config.Password != "" {
		credential = fmt.Sprintf("%s:%s@", config.Username, config.Password)
	}
	return fmt.Sprintf("mongodb://%s%s/%s?maxPoolSize=%d", credential,
		strings.Join(config.Address, ","), config.Database, config.MaxPoolSize)
}

func shouldRetry(ctx context.Context, err error) bool {
	select {
	case <-ctx.Done():
		return false
	default:
		if cmdErr, ok := err.(mongo.CommandError); ok {
			return cmdErr.Code != 13
		}
		return true
	}
}
