package modules

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"
)

var AppCache cache.Cache

func init() {
	beego.Info("init cache.")
	c, _ := cache.NewCache("memory", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120,"interval":60}`)
	AppCache = c
}

func JsonRequestCache(ctx *context.Context, key string) bool {
	ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")

	cacheData := AppCache.Get(key)
	if cacheData == nil {
		return false
	}

	var str string = cacheData.(string)
	ctx.Output.Body([]byte(str))

	return true
}