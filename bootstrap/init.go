package bootstrap

import (
	"github.com/YogeLiu/CloudDisk/dao"
	"github.com/YogeLiu/CloudDisk/pkg/cache"
	"github.com/YogeLiu/CloudDisk/pkg/conf"
	"github.com/gin-gonic/gin"
)

func Init(path string) {
	conf.Init(path)
	if conf.SystemConfig.Debug {
		gin.SetMode(gin.TestMode)
	}
	module := []struct{ factory func() }{
		// databse
		{dao.Init},
		// cache
		{cache.Init},
	}

	for _, item := range module {
		item.factory()
	}
}
