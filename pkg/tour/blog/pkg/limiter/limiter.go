package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type Iface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBucket(rules ...BucketRule) Iface
}

type BucketRule struct {
	Key          string
	FillInternal time.Duration
	Capacity     int64
	Quantum      int64
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}
