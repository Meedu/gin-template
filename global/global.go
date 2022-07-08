package global

import (
	"sync"

	"github.com/Meedu/gin-template/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/Meedu/gin-template/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MD_DB     *gorm.DB
	MD_DBList map[string]*gorm.DB
	MD_REDIS  *redis.Client
	MD_CONFIG config.Server
	MD_VP     *viper.Viper
	// MD_LOG    *oplogging.Logger
	MD_LOG                 *zap.Logger
	MD_Timer               timer.Timer = timer.NewTimerTask()
	MD_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)
